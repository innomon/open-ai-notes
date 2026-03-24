---
name: frappe-expert
description: Expert guidance for Frappe Framework development, including setup, app creation, site management, local testing, and production deployment. Use when creating or managing Frappe-based applications and ERPNext extensions.
---

# Frappe Expert Skill

This skill transforms Gemini CLI into a senior Frappe/ERPNext developer. It provides structured workflows for setting up the environment, creating new apps, and deploying them to production.

## Core Workflows

### 1. Initial Environment Setup
When asked to set up a new Frappe environment:
1.  **Check Prerequisites**: Ensure Python 3.10+, Node.js 18+, MariaDB 10.6+, and Redis are installed.
2.  **Install Bench**: `pip3 install frappe-bench`
3.  **Initialize Bench**: `bench init <bench_name>`
4.  **Create Site**: `bench new-site <site_name>`

### 2. Creating a New Frappe App
When given a high-level specification for a new app:
1.  **Generate Specification**: Create a `docs/specification.md` based on [spec-template.md](references/spec-template.md).
2.  **Generate Implementation Plan**: Create a `docs/implementation_plan.md` with a phased checklist using [checklist.md](references/checklist.md).
3.  **Initialize App**: `bench new-app <app_name>`
4.  **Install App**: `bench --site <site_name> install-app <app_name>`
5.  **Create agents.md**: Initialize an `agents.md` in the project root using [agents-template.md](references/agents-template.md) to ensure persistent agent context.

### 3. Local Development & Testing
- **Start Services**: `bench start` (to run the dev server)
- **Run Tests**: `bench --site <site_name> run-tests --app <app_name>`
- **Migrations**: Always run `bench --site <site_name> migrate` after DocType changes.

### 4. Cloud Deployment
For cloud deployment, prioritize these methods:
- **Frappe Cloud**: Recommended for managed hosting.
- **Docker (Frappe Docker)**: For containerized environments.
- **Manual VPS**: See [deployment.md](references/deployment.md) for Supervisor/Nginx configuration.

## Design Patterns & Standards
- **DocTypes**: Use "Standard" DocTypes unless "Virtual" is required.
- **Naming**: Use `snake_case` for fields and `PascalCase` for DocTypes.
- **Logic**: Keep business logic in the DocType class (Controller) or in separate service modules.
- **UI**: Use Frappe Desk for admin UI; custom Web Pages for public UI.

## Troubleshooting
- **Database Connection**: Verify MariaDB is running and accessible.
- **Permissions**: Ensure the user has `sudo` access for production setup.
- **Port Conflicts**: Default bench port is 8000; default MariaDB is 3306.
