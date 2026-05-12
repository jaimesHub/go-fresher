# Project Setup Guide

This document describes the directory structure and setup process for the Go Fresher backend project.

## Go Module Initialization

```bash
go mod init github.com/jaimesHub/go-fresher
# Similar to: npm init
```

## Directory Structure

### `cmd/` - Application Entry Points
Executable entry points for the application. Each subdirectory is a standalone binary.

- **server/** - HTTP API server
  - `main.go` - Entry point
  - Run with: `go run cmd/server/main.go`
  
- **cronjob/** - Scheduled background jobs and workers
  
- **cli/** - Command-line interface tools

### `internal/` - Private Application Code
Internal packages used only by this project. Not exported to other modules.

- **controller/** - HTTP request handlers and API endpoints
- **models/** - Domain models and data structures
- **service/** - Business logic layer (core functionality)
- **repo/** - Data access layer (Repository/DAO pattern)
- **router/** - HTTP route definitions and setup
- **middlewares/** - HTTP middleware (auth, logging, error handling, etc.)
- **initialize/** - Application initialization and startup logic

### `pkg/` - Shared Utilities
Reusable packages that can be shared across multiple projects in the `/go` directory.

- **utils/** - Utility helper functions
- **setting/** - Configuration management
- **logger/** - Logging utilities

### Other Directories

- **configs/** - Configuration files (database, environment settings, etc.)
- **migrations/** - Database migration scripts
  - `1_create_table.up.sql` - Schema creation SQL
- **scripts/** - Build scripts and utility scripts
- **tests/** - Test files and test utilities
- **docs/** - Internal documentation for developers
  - Note: Add to `.gitignore` if contains sensitive info
- **third_party/** - External dependencies and third-party code
- **response/** - HTTP response utilities and helpers
  - `httpStatusCode.go` - HTTP status code definitions
- **global/** - Global constants and application-wide variables

## Getting Started

1. Initialize the Go module (if not already done)
2. Ensure all directories are created
3. Start implementing handlers in `internal/controller/`
4. Define routes in `internal/router/`
5. Write business logic in `internal/service/`
6. Implement data access in `internal/repo/`
7. Add shared utilities to `pkg/`

## Running the Application

```bash
# Start the server
go run cmd/server/main.go

# Run background jobs
go run cmd/cronjob/main.go

# Use CLI tools
go run cmd/cli/main.go [command]
```