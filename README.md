# Go Fresher - Backend Project

A Go backend project template with a clean architecture following standard Go project layout conventions.

## 📋 Project Overview

**Go Fresher** is a backend application built with Go, designed with a scalable and maintainable architecture. It includes structured patterns for HTTP servers, background jobs, CLI tools, and database migrations.

## 🗂️ Project Structure

```
go-fresher/
├── cmd/                    # Executable entry points
│   ├── server/            # HTTP API server
│   ├── cronjob/           # Scheduled background jobs
│   └── cli/               # Command-line tools
├── internal/              # Private application code
│   ├── controller/        # HTTP request handlers
│   ├── models/            # Data structures/domain models
│   ├── service/           # Business logic layer
│   ├── repo/              # Data access layer (DAO)
│   ├── routers/            # HTTP route definitions
│   ├── middlewares/       # HTTP middleware handlers
│   └── initialize/        # Application initialization
├── pkg/                   # Reusable packages (shared across projects)
│   ├── utils/             # Utility functions
│   ├── setting/           # Configuration management
│   └── logger/            # Logging utilities
├── configs/               # Configuration files
├── migrations/            # Database migrations
│   └── 1_create_table.up.sql
├── scripts/               # Build and utility scripts
├── tests/                 # Test files
├── docs/                  # Project documentation
├── third_party/           # External dependencies/libraries
├── response/              # HTTP response utilities
│   └── httpStatusCode.go
├── global/                # Global constants and variables
└── go.mod                 # Go module definition
```

## 🚀 Quick Start

### Prerequisites
- Go 1.16+ installed
- Database setup (if needed)

### Running the Server

```bash
go run cmd/server/main.go
```

### Running Cronjobs

```bash
go run cmd/cronjob/main.go
```

### CLI Tools

```bash
go run cmd/cli/main.go [command] [options]
```

## 🏗️ Architecture

The project follows **Clean Architecture** principles:

1. **Controllers** - Handle HTTP requests/responses
2. **Services** - Contain business logic
3. **Repository** - Abstract data access (DAO pattern)
4. **Models** - Domain models and data structures
5. **Router** - Define API routes and endpoints
6. **Middlewares** - Cross-cutting concerns (auth, logging, etc.)

## 📦 Configuration

Configuration files are stored in the `configs/` directory. The application uses the `pkg/setting` package to manage configurations.

## 🔄 Database Migrations

Database migrations are stored in the `migrations/` directory. Use a migration tool like `golang-migrate` to manage schema changes.

```bash
# Example migration command
migrate -path migrations -database "your-db-url" up
```

## 📝 Development

- Write business logic in `internal/service/`
- Add HTTP handlers in `internal/controller/`
- Implement data access in `internal/repo/`
- Add reusable utilities in `pkg/`
- Update routes in `internal/router/`

## 📚 Documentation

Additional documentation can be found in the `docs/` directory.

## 📄 License

Add your license information here.
