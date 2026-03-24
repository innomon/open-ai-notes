---
name: gcloud-deploy
description: Deploys the Go/Next.js project to Google Cloud Run using the custom deployment utility.
version: 1.0.0
author: Gemini
tags: [deploy, gcp, cloud-run, golang]
---

# Google Cloud Run Deployment Skill

## Description
This skill provides the standard operating procedure for an AI agent to deploy the current project to Google Cloud Run using the custom Go-based deployment utility located at `cmd/deploy/main.go`.

## Context
Invoke this skill whenever the user asks to "deploy the app", "push to production", "deploy to cloud run", or specifically requests to run the deployment script. The deployment script packages the local source, uploads it to Google Cloud Storage, triggers Google Cloud Build to create a Docker container, and finally deploys that container to Google Cloud Run.

## Prerequisites
Before initiating the deployment, verify the following:
1. **Authentication:** The user must be authenticated with Google Cloud CLI (`gcloud auth login` or default application credentials).
2. **Environment Variables:** The `.env` or `.env.local` file must contain at least `GOOGLE_CLOUD_PROJECT` (or `PROJECT_ID`, or `NEXT_PUBLIC_FIREBASE_PROJECT_ID`).
3. **Dependencies:** The deployment script requires Go to be installed on the system.

## Instructions

1. **Verify Environment Configuration:**
   - Check if `GOOGLE_CLOUD_PROJECT` or a fallback project ID exists in `.env` or `.env.local`.
   - If not found, prompt the user to provide the target Google Cloud Project ID before proceeding.

2. **Execute the Deployment Utility:**
   - Run the following command from the root of the project:
     ```bash
     go run cmd/deploy/main.go
     ```
   - *Note:* Do not build the binary first unless `go run` fails. `go run` is sufficient for execution.

3. **Monitor the Execution:**
   - The script will output several stages:
     - "Creating source archive..."
     - "Uploading source to GCS..."
     - "Triggering Cloud Build..."
     - "Deploying to Cloud Run..."
   - Wait for the process to complete. It may take several minutes during the Cloud Build phase.

4. **Post-Deployment:**
   - Upon seeing "Deployment successful!" and the Cloud Run service URL, report the successful deployment to the user and provide them with the deployed URL.
   - Remind the user to do a "Hard Refresh" in their browser to clear any cached static assets if they were experiencing stale UI issues prior to the deployment.

5. **Error Handling:**
   - If the script fails at the "Uploading source to GCS..." stage, ensure the user has the correct GCP permissions (Storage Admin).
   - If Cloud Build fails, suggest checking the Google Cloud Console for detailed build logs.
   - If the deployment to Cloud Run fails, ensure the Cloud Run Admin API is enabled on the GCP project.
