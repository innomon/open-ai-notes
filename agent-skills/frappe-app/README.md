# Frappe Expert Skill & Agent Configuration

This repository provides a specialized **Frappe Expert Skill** for Gemini CLI and a standard **agents.md** configuration for AI coding agents (Claude Code, Amp Code, etc.).

## 🚀 Skill Installation (Gemini CLI)

The `frappe-expert.skill` file transforms Gemini CLI into a senior Frappe developer with deep knowledge of `bench`, DocTypes, and Frappe deployment patterns.

### Step 1: Install the Skill
Choose your preferred scope:

- **User Level** (Available across all your projects):
  ```bash
  gemini skills install frappe-expert.skill --scope user
  ```
- **Workspace Level** (Available only in this project):
  ```bash
  gemini skills install frappe-expert.skill --scope workspace
  ```

### Step 2: Reload Gemini CLI
After installation, you **must** reload the skills in your active Gemini CLI session:
```bash
/skills reload
```

---

## 🤖 Agent Configuration (Claude Code, Amp Code, etc.)

For agents that support the [agents.md](https://agents.md/) specification (like Claude Code and Amp Code), you should place an `agents.md` file in your project root.

### Recommended `agents.md` for Frappe
Create a file named `agents.md` (or `AGENTS.md`) in your repository root with the following content to give AI agents immediate context:

```markdown
# Project: [Your App Name]
This is a Frappe-based application.

## Tech Stack
- Framework: Frappe Framework (Python/JS/HTML)
- Tooling: Bench CLI
- Database: MariaDB (standard)

## Development Workflow
- Start Server: `bench start`
- New App: `bench new-app <app_name>`
- Create Site: `bench new-site <site_name>`
- Run Tests: `bench --site <site_name> run-tests --app <app_name>`

## Coding Rules
- Use `snake_case` for field names; `PascalCase` for DocTypes.
- Always use `@frappe.whitelist()` for web-accessible functions.
- Business logic belongs in the DocType's Python controller.
```

---

## 🛠 Usage & Workflow

Once the skill is active (or `agents.md` is present), you can use high-level prompts like:

> "Set up a new Frappe environment and create a Library Management app."
>
> "Create a specification and implementation plan for a new Frappe app that manages employee attendance using geolocation."
>
> "Deploy this app to a VPS using Supervisor and Nginx."

### Core Commands Reference
| Command | Action |
| :--- | :--- |
| `bench init [name]` | Initialize a new Frappe environment |
| `bench new-app [app]` | Create a new application |
| `bench new-site [site]` | Create a new database instance |
| `bench start` | Start the local development server |
| `bench run-tests` | Execute the unit test suite |
