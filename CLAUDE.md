# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based CLI tool for interacting with Bitbucket Enterprise REST APIs. The tool provides functionality for managing projects, repositories, and pull requests.

## Build and Development Commands

**Build the application:**
```bash
just build
```
This creates a binary at `./build/bitbucket-cli`

**Clean build artifacts:**
```bash
just clean
```

**Run tests:**
```bash
go test ./...
```
Note: Tests require proper Bitbucket API configuration and may fail without valid credentials/URLs.

**Run a single test:**
```bash
go test -run TestName ./internal
```

## Architecture

The application follows a command-line architecture with the following key components:

### Main Entry Point
- `cmd/bitbucket-cli/bitbucket.go` - Main application entry point with argument parsing and configuration loading

### Core Components
- `internal/cli.go` - Core CLI struct and BitbucketCLI client initialization
- `internal/auth.go` - Authentication handling (Basic Auth and Access Token)
- `internal/client.go` - Bitbucket API client wrapper
- `internal/config.go` - Configuration management with YAML support

### Command Modules
- `internal/project.go` + `internal/project_*.go` - Project-related commands (list, clone)
- `internal/repo.go` + `internal/repo_*.go` - Repository-related commands
- `internal/pr.go` + `internal/pr_*.go` - Pull request commands (create, list, approve, merge)

### Key Dependencies
- `github.com/gfleury/go-bitbucket-v1` - Bitbucket API client
- `github.com/alexflint/go-arg` - Command-line argument parsing
- `github.com/go-git/go-git/v5` - Git operations for cloning
- `github.com/sirupsen/logrus` - Logging

## Configuration

The CLI supports configuration via:
1. Command-line arguments
2. Environment variables (BITBUCKET_USERNAME, BITBUCKET_PASSWORD, BITBUCKET_ACCESS_TOKEN, BITBUCKET_URL)
3. YAML config file at `~/.config/bitbucket-cli/config.yml`

Required configuration:
- Bitbucket REST API URL (--url or BITBUCKET_URL)
- Username (--username or BITBUCKET_USERNAME)
- Either password (--password) or access token (--access-token)

## Command Structure

The CLI follows a hierarchical command structure:
- `project` - Project operations (list, clone)
- `repo` - Repository operations (requires -k KEY -n NAME)
  - `pr` - Pull request operations (create, list, approve, merge)
  - `branch` - Branch operations
  - `security` - Security scanning