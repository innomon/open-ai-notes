# Project: <App Name>
This is a Frappe-based application.

## Tech Stack
- **Framework**: Frappe Framework (Python/JS/HTML)
- **Database**: MariaDB (standard) or PostgreSQL
- **Environment**: Frappe Bench

## Development Workflow
- **Start Dev Server**: `bench start`
- **Run Tests**: `bench --site <site_name> run-tests --app <app_name>`
- **Migrations**: `bench --site <site_name> migrate`
- **DocType Creation**: Create DocTypes in Frappe Desk and export to app.

## Rules & Conventions
- **Naming**: Use `snake_case` for field names and `PascalCase` for DocTypes.
- **API**: Use `@frappe.whitelist()` for any function accessible via the web.
- **Permissions**: Always check permissions in the controller logic using `frappe.has_permission`.
- **UI**: Prefer standard Frappe Desk UI for administrative tasks.

## Key Files
- `hooks.py`: Main configuration and event hooks.
- `public/js/`: Client-side scripts.
- `<app_name>/<module_name>/doctype/`: DocType definitions and controllers.

## Deployment
- **Method**: <e.g., Frappe Cloud, Docker, Manual>
- **Production Site**: <site_url>
