<p align="center">
  <a href="https://indent.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://indent.com/static/indent_text_white.png">
      <img src="https://indent.com/static/indent_text_black.png" height="48">
    </picture>
    <h1 align="center"><code>access</code></h1>
    <p align="center">Easiest way to request and grant<br />access without leaving your terminal</p>
  </a>
</p>

<p align="center">
  <a aria-label="Made by Indent" href="https://indent.com/">
    <img alt="Made by Indent" src="https://img.shields.io/badge/Made%20By%20Indent-fff.svg?style=flat&labelColor=2f80ed">
  </a>
  <a aria-label="License" href="https://github.com/vercel/turbo/blob/main/LICENSE">
    <img alt="Badge for Apache 2 License" src="https://img.shields.io/badge/license-Apache%202-blue">
  </a>
  <a aria-label="GitHub Workflow " href="https://indent.com/">
    <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/indentapis/access/build-access.yaml">
  </a>
</p>

## What is `access`?

`access` is a command-line tool for [Indent](https://indent.com) that allows users to request access to apps and services across their organization. This provides time-bound access requests so users only have permissions when they need them while staying productive.

You can request or manage access directly from the terminal:

- `access petitions create --resources=<id> --reason=<reason>`
- `access petitions approve <petition> --duration=6h`
- `access petitions revoke <petition>`

## Getting Started

Visit https://access.new to request access on the web, or https://indent.com/setup to set up an account.

## Quick Start

If you're on Mac, you can install via Homebrew: _(in progress Homebrew PR)_

```bash
brew tap indentapis/homebrew-access
brew install access
```

Or to download an `access` release directly, follow these steps:

1. Download the latest binary from [releases](https://github.com/indentapis/access/releases).
2. Run `access init <space>` command to log in and set up your configuration.
   * Replace `<space>` with your space name from [indent.com/spaces](https://indent.com/spaces)
3. Use the available commands to request and review access, or type `access --help` for options.

## Commands

`access` provides several commands to perform various operations related to authentication, configuration, and access requests.

Here is a summary of the available commands:

| Command                    | Description                                                 |
|----------------------------|-------------------------------------------------------------|
| `access auth`              | Perform operations related to Indent authentication.        |
| `access completion`        | Generate the autocompletion script for the specified shell. |
| `access config`            | Make changes to the access configuration.                   |
| `access help`              | Get help about any access command.                          |
| `access init`              | Set up access for first-time use.                           |
| `access petitions approve` | Approve a Petition for a specified amount of time.          |
| `access petitions close`   | Close a Petition.                                           |
| `access petitions create`  | Request access to a Resource                                |
| `access petitions deny`    | Deny a Petition.                                            |
| `access petitions list`    | List all the Petitions.                                     |
| `access petitions revoke`  | Revoke a Petition and related access.                       |
| `access resources`         | Manage Resources within a space.                            |

For more information about each command, run `access [command] --help`.
