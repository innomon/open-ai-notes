# GCloud Deploy Skill

This directory contains the `SKILL.md` specification file for deploying the Go/Next.js application to Google Cloud Run using the `cmd/deploy/main.go` deployment utility. 

It is designed following the [agents.md](https://agents.md/) specification to provide clear, actionable instructions to AI coding agents, ensuring they follow your project's specific deployment pipeline rather than inventing their own.

## Installation & Usage

### 1. Gemini CLI
The Gemini CLI natively supports skills. You can load this skill directly when running Gemini or inside an interactive session.
- **CLI Flag:** `gemini --skill ./gcloud-deploy-skill/SKILL.md`
- **In-Chat Instruction:** Just ask Gemini: "Activate the gcloud-deploy skill from the gcloud-deploy-skill directory and deploy the app."

### 2. Claude Code
Claude Code looks for a `CLAUDE.md` file in the root of your project to understand repository-specific context and commands.
- **Option 1 (Reference):** Add a note to your root `CLAUDE.md` pointing to this directory:
  ```markdown
  ## Deployment
  To deploy this application, strictly follow the instructions defined in `gcloud-deploy-skill/SKILL.md`.
  ```
- **Option 2 (Direct Prompt):** Tell Claude Code: "Deploy the app using the instructions found in `gcloud-deploy-skill/SKILL.md`."

### 3. AMP Code (Aider / Cursor / Windsurf)
Most modern AI coding tools (AMP Code, Aider, Cursor, Roo Code) support adding specific context files to the chat or using an `.md` file for project rules.
- **Cursor/Windsurf:** You can either `@` reference the `gcloud-deploy-skill/SKILL.md` file in the chat and say "Deploy the app using these instructions," or add its contents to your `.cursorrules` / `.windsurfrules` file.
- **Aider / Roo Code:** Start the agent and explicitly include the skill file in the context window:
  ```bash
  aider --message "Deploy the app based on the gcloud-deploy-skill/SKILL.md instructions."
  ```

## What the Skill Does
When activated, the AI agent will:
1. Verify GCP project configurations in your `.env` files.
2. Execute the `cmd/deploy/main.go` utility.
3. Monitor the Cloud Build, GCS upload, and Cloud Run deployment steps.
4. Notify you when the deployment completes successfully with the live URL.