# Frappe Deployment Guide

## 1. Managed Hosting (Frappe Cloud)
- **Recommended for**: Ease of use, scalability, and built-in backups.
- **Workflow**:
  1. Push app to a public or private GitHub/GitLab repository.
  2. Create a new "Bench" on [Frappe Cloud](https://frappecloud.com).
  3. Connect the repository and install the app.
  4. Deploy to a site.

## 2. Containerized Deployment (Docker)
- **Recommended for**: Multi-app isolation and local parity.
- **Resources**: [Frappe Docker GitHub](https://github.com/frappe/frappe-docker)
- **Workflow**:
  1. Use `docker-compose.yml` to define services (Redis, DB, Web, Worker).
  2. Build custom images including your apps.
  3. Deploy using Docker Swarm or Kubernetes (Helm charts available).

## 3. Manual Production Setup (VPS)
- **Recommended for**: Maximum control and custom server environments.
- **Easy Install (Ubuntu/Debian)**:
  ```bash
  wget https://raw.githubusercontent.com/frappe/bench/develop/easy-install.py
  sudo python3 easy-install.py --production --user <username>
  ```
- **Manual Setup (Supervisor & Nginx)**:
  1. Initialize production setup: `sudo bench setup production <username>`
  2. Generate Nginx config: `bench setup nginx`
  3. Generate Supervisor config: `bench setup supervisor`
  4. Symlink configs:
     ```bash
     sudo ln -s `pwd`/config/nginx.conf /etc/nginx/conf.d/<site_name>.conf
     sudo ln -s `pwd`/config/supervisor.conf /etc/supervisor/conf.d/<site_name>.conf
     ```
  5. Reload services:
     ```bash
     sudo service nginx reload
     sudo supervisorctl reload
     ```

## 4. Maintenance Commands
- **Update Bench and Apps**: `bench update`
- **Run Migrations**: `bench --site <site_name> migrate`
- **Backup Site**: `bench --site <site_name> backup`
- **Check Logs**: `bench --site <site_name> logs` or `tail -f logs/*.log`
