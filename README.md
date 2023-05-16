<p align="center">
  <a href="https://indent.com">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://indent.com/static/indent_text_white.png">
      <img src="https://indent.com/static/indent_text_black.png" height="32">
    </picture>
  </a>
  <h1 align="center"><code>access</code></h1>
  <p align="center"><code>$ <code>brew install access</code></code></p>
  <p align="center">Easiest way to request and grant<br />access without leaving your terminal</p>
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

`access` is an all-in-one command-line tool from [Indent](https://indent.com) for requesting and managing temporary access for cloud apps and systems.

This provides users with temporary access that's fast and easy to get because it's on-demand and automated, so users only have permissions when they need them without having to wait hours or days to get access.

```sh
access new             # Search for access, provide a reason, submit request
```

The `access` command-line tool works with multiple applications, identity providers, and cloud infrastucture providers. Instead of referencing a bunch of different docs or who's the admin for some service, you can just type `access`. Indent will automatically route requests submitted by `access` to the right reviewer, they can approve directly from Slack or the command-line, then you'll get a notification that your access has been granted.

```sh
alias prodlogs="access petitions create --resources=e843ad66"

$ prodlogs --reason "to debug INC-4881"
$ access petitions approve 42010a1c010b --duration=6h
$ access petitions revoke 42010a1c010
```

## Getting Started

Type [access.new](https://access.new) to request access on the web, or go to [indent.com/setup](https://indent.com/setup) to set up an account.

If you're on Mac, you can install via Homebrew:

```bash
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
