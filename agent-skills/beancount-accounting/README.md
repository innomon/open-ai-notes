# Beancount Accounting Skill

A specialized skill for Gemini CLI and other Model Context Protocol (MCP) compatible agents to handle end-to-end financial accounting using the [Beancount](https://beancount.github.io/) plain-text accounting language.

## 🌟 Overview

This skill transforms a general-purpose AI agent into a senior accountant capable of processing raw documents (PDF/CSV bank statements, invoices) into a structured double-entry ledger. It provides automated workflows for reconciliation, financial reporting, and seamless data migration to popular accounting platforms.

## 🛠 Features

- **Automated Data Ingestion:** Strategies for extracting transaction data from PDF/CSV bank statements.
- **Beancount Lifecycle Management:** Tools to initialize workspaces, maintain a hierarchical Chart of Accounts, and verify ledger integrity.
- **Multi-Format Financial Exports:**
  - **Tally ERP 9 / TallyPrime:** Generates compliant XML for Masters and Vouchers.
  - **Zoho Books:** Generates CSV for Manual Journal imports.
  - **Orez Books / Frappe Books:** Generates CSV optimized for Orez/Frappe Books journal entry imports.
  - **Generic CSV:** Standard flat-file export for custom analysis.
- **Reporting:** Real-time generation of Profit & Loss (P&L) statements and Balance Sheets.

## 📁 Structure

```text
beancount-accounting/
├── SKILL.md            # Core instructions and triggering metadata
├── scripts/
│   └── main.py         # Consolidated CLI tool for init, report, and export
└── references/
    ├── beancount_syntax.md # Quick reference for Beancount language
    └── formats.md          # Technical specifications for export formats
```

## 🚀 Usage

### 1. Installation

If you have the `.skill` package:
```bash
gemini skills install beancount-accounting.skill --scope user
/skills reload
```

### 2. Initializing a Workspace
The agent can bootstrap a standard accounting folder structure:
```bash
python3 scripts/main.py init
```

### 3. Workflow Commands

- **Verify & Report:**
  ```bash
  python3 scripts/main.py report
  ```
- **Export for Tally:**
  ```bash
  python3 scripts/main.py export --format tally
  ```
- **Export for Orez/Frappe Books:**
  ```bash
  python3 scripts/main.py export --format orez
  ```
- **Export for Zoho Books:**
  ```bash
  python3 scripts/main.py export --format zoho
  ```

## 🤖 AI Agent Integration

This skill is designed to be used by agents that support the **Gemini CLI Skill System**. Once activated, the agent will automatically:
1. Identify bank transactions and map them to the correct accounts.
2. Ensure the ledger balances correctly (Assets = Liabilities + Equity).
3. Suggest the correct export format based on your destination accounting software.

## 📜 License

[AGPL-3.0](LICENSE) (Matches Orez/Frappe Books foundation)
