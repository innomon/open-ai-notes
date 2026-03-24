package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1/v2"
	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Identify project root and load environment variables
	root, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Could not find project root: %v", err)
	}
	fmt.Printf("Project root: %s\n", root)

	// Change working directory to root for easier path handling
	if err := os.Chdir(root); err != nil {
		log.Fatalf("Failed to change to root directory: %v", err)
	}

	godotenv.Load(".env")
	godotenv.Load(".env.local")
	// godotenv.Load("../../.env")
	// godotenv.Load("../../.env.local")

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = os.Getenv("PROJECT_ID")
	}
	if projectID == "" {
		projectID = os.Getenv("NEXT_PUBLIC_FIREBASE_PROJECT_ID")
	}
	if projectID == "" {
		log.Fatal("GOOGLE_CLOUD_PROJECT, PROJECT_ID or NEXT_PUBLIC_FIREBASE_PROJECT_ID environment variable is required")
	}

	region := os.Getenv("GOOGLE_CLOUD_REGION")
	if region == "" {
		region = "us-central1"
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "farmer-kheti-badi"
	}

	ctx := context.Background()

	// 2. Prepare source archive
	fmt.Println("Creating source archive...")
	archivePath := filepath.Join(os.TempDir(), fmt.Sprintf("source-%d.tar.gz", time.Now().Unix()))
	if err := createArchive(archivePath); err != nil {
		log.Fatalf("Failed to create archive: %v", err)
	}
	defer os.Remove(archivePath)

	// 3. Upload to GCS
	fmt.Println("Uploading source to GCS...")
	bucketName := fmt.Sprintf("%s_cloudbuild_source", projectID)
	objectName := fmt.Sprintf("source-%d.tar.gz", time.Now().Unix())

	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create storage client: %v", err)
	}
	defer storageClient.Close()

	bucket := storageClient.Bucket(bucketName)
	if _, err := bucket.Attrs(ctx); err != nil {
		fmt.Printf("Bucket %s does not exist, creating it...\n", bucketName)
		if err := bucket.Create(ctx, projectID, nil); err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	}

	obj := bucket.Object(objectName)
	w := obj.NewWriter(ctx)
	f, err := os.Open(archivePath)
	if err != nil {
		log.Fatalf("Failed to open archive: %v", err)
	}
	if _, err := io.Copy(w, f); err != nil {
		log.Fatalf("Failed to upload to GCS: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close GCS writer: %v", err)
	}
	f.Close()
	fmt.Printf("Uploaded source to gs://%s/%s\n", bucketName, objectName)

	// 4. Trigger Cloud Build
	fmt.Println("Triggering Cloud Build...")
	cbClient, err := cloudbuild.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create cloudbuild client: %v", err)
	}
	defer cbClient.Close()

	timestamp := time.Now().Unix()
	imageName := fmt.Sprintf("gcr.io/%s/%s:%d", projectID, serviceName, timestamp)

	buildArgs := map[string]string{}
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if strings.HasPrefix(pair[0], "NEXT_PUBLIC_") {
			buildArgs[pair[0]] = pair[1]
		}
	}

	dockerArgs := []string{"build", "--tag", imageName}
	for k, v := range buildArgs {
		dockerArgs = append(dockerArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}
	dockerArgs = append(dockerArgs, ".")

	buildReq := &cloudbuildpb.CreateBuildRequest{
		ProjectId: projectID,
		Build: &cloudbuildpb.Build{
			Source: &cloudbuildpb.Source{
				Source: &cloudbuildpb.Source_StorageSource{
					StorageSource: &cloudbuildpb.StorageSource{
						Bucket: bucketName,
						Object: objectName,
					},
				},
			},
			Steps: []*cloudbuildpb.BuildStep{
				{
					Name: "gcr.io/cloud-builders/docker",
					Args: dockerArgs,
				},
			},
			Images: []string{imageName},
		},
	}

	op, err := cbClient.CreateBuild(ctx, buildReq)
	if err != nil {
		log.Fatalf("Failed to create build: %v", err)
	}

	fmt.Printf("Build started: %s\n", op.Name())

	// Wait for build
	finishedBuild, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("Build failed: %v", err)
	}
	if finishedBuild.Status != cloudbuildpb.Build_SUCCESS {
		log.Fatalf("Build finished with status: %s", finishedBuild.Status)
	}
	fmt.Println("Build finished successfully.")

	// 5. Deploy to Cloud Run
	fmt.Println("Deploying to Cloud Run...")
	runClient, err := run.NewServicesClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create run client: %v", err)
	}
	defer runClient.Close()

	runEnv := []*runpb.EnvVar{
		{
			Name: "GOOGLE_CLOUD_PROJECT",
			Values: &runpb.EnvVar_Value{
				Value: projectID,
			},
		},
	}
	envMap, _ := godotenv.Read(".env", ".env.local")
	for k, v := range envMap {
		if !strings.HasPrefix(k, "NEXT_PUBLIC_") && k != "GOOGLE_CLOUD_PROJECT" {
			runEnv = append(runEnv, &runpb.EnvVar{
				Name: k,
				Values: &runpb.EnvVar_Value{
					Value: v,
				},
			})
		}
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", projectID, region)
	serviceID := serviceName
	servicePath := fmt.Sprintf("%s/services/%s", parent, serviceID)

	serviceConfig := &runpb.Service{
		Name: servicePath,
		Template: &runpb.RevisionTemplate{
			Containers: []*runpb.Container{
				{
					Image: imageName,
					Env:   runEnv,
					Ports: []*runpb.ContainerPort{
						{
							ContainerPort: 8080,
						},
					},
				},
			},
		},
	}

	// Check if service exists
	existingService, err := runClient.GetService(ctx, &runpb.GetServiceRequest{Name: servicePath})
	if err != nil {
		fmt.Printf("Service %s not found, creating...\n", serviceID)
		opCreate, err := runClient.CreateService(ctx, &runpb.CreateServiceRequest{
			Parent:    parent,
			Service:   serviceConfig,
			ServiceId: serviceID,
		})
		if err != nil {
			log.Fatalf("Failed to create service: %v", err)
		}
		res, err := opCreate.Wait(ctx)
		if err != nil {
			log.Fatalf("Service creation failed: %v", err)
		}
		fmt.Printf("Service created: %s\n", res.Uri)
	} else {
		fmt.Printf("Service %s found, updating...\n", serviceID)
		// Maintain existing properties if needed, but here we overwrite with new config
		serviceConfig.Etag = existingService.Etag
		opUpdate, err := runClient.UpdateService(ctx, &runpb.UpdateServiceRequest{
			Service: serviceConfig,
		})
		if err != nil {
			log.Fatalf("Failed to update service: %v", err)
		}
		res, err := opUpdate.Wait(ctx)
		if err != nil {
			log.Fatalf("Service update failed: %v", err)
		}
		fmt.Printf("Service updated: %s\n", res.Uri)
	}

	fmt.Println("Deployment successful!")
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// Check if it's the root go.mod or backend/go.mod
			if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
				return dir, nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	// Fallback: check current directory and then parent if we are in cmd/deploy
	cwd, _ := os.Getwd()
	if strings.HasSuffix(cwd, filepath.Join("cmd", "deploy")) {
		return filepath.Dir(filepath.Dir(cwd)), nil
	}
	return cwd, nil
}

func createArchive(target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Ignore list
	ignore := map[string]bool{
		"node_modules":  true,
		".git":          true,
		".next":         true,
		"out":           true,
		"deploy_tool":   true,
		"source.tar.gz": true,
	}

	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip ignored directories
		parts := strings.Split(path, string(os.PathSeparator))
		for _, part := range parts {
			if ignore[part] {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = path

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tw, file)
		return err
	})
}
