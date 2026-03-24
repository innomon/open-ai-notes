# Frappe App Development Checklist

## Phase 1: Planning & Setup
- [ ] Initialize Frappe Bench and Site
- [ ] Create New App: `bench new-app <app_name>`
- [ ] Install App on Site: `bench --site <site_name> install-app <app_name>`
- [ ] Create initial `agents.md` in project root

## Phase 2: Core Data Structure (DocTypes)
- [ ] Define DocTypes and Fields (JSON/Python)
- [ ] Set Naming Conventions for DocTypes
- [ ] Configure Permissions (Role Permissions Manager)
- [ ] Run `bench --site <site_name> migrate`

## Phase 3: Business Logic
- [ ] Implement Controller Hooks (before_insert, on_update, validate)
- [ ] Create Custom API Endpoints (whitelisted functions)
- [ ] Add Form Scripts (Client-side logic)
- [ ] Define Hooks in `hooks.py` (e.g., scheduled jobs, event handlers)

## Phase 4: User Interface
- [ ] Customize List Views and Form Layouts
- [ ] Create Dashboards and Reports (Script Report, Query Report)
- [ ] Design Web Pages/Portal Views (if applicable)

## Phase 5: Testing & Quality Assurance
- [ ] Write Unit Tests (`tests/test_<doctype>.py`)
- [ ] Run Tests: `bench --site <site_name> run-tests --app <app_name>`
- [ ] Perform UI testing (Frappe Desk)
- [ ] Verify Role-based Permissions

## Phase 6: Deployment & CI/CD
- [ ] Configure `supervisor.conf` and `nginx.conf` (Production)
- [ ] Set up Periodic Backups
- [ ] Run `bench build` for production assets
- [ ] Verify app works in production environment
