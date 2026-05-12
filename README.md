# Go Fresher - Backend Project

A Go backend project built with **Gin Gonic** and clean architecture principles following standard Go project layout conventions.

## 📋 Project Overview

**Go Fresher** is a learning project demonstrating a scalable Go backend architecture. It showcases HTTP API development with Gin, clean architecture patterns, API versioning, and structured request handling.

**Current Status:** Working HTTP API server with versioned endpoints (v1, v2)

## 🗂️ Project Structure

```
go-fresher/
├── cmd/                      # Executable entry points
│   ├── server/              # HTTP API server (main application)
│   ├── cronjob/             # Scheduled background jobs
│   └── cli/                 # Command-line tools
├── internal/                # Private application code
│   ├── routers/             # HTTP route definitions & handlers
│   ├── controller/          # HTTP request handlers (ready for expansion)
│   ├── service/             # Business logic layer (ready for expansion)
│   ├── models/              # Data structures & domain models
│   ├── repo/                # Data access layer (DAO pattern)
│   ├── middlewares/         # HTTP middleware (auth, logging, etc.)
│   └── initialize/          # Application startup logic
├── pkg/                     # Reusable packages (shareable across projects)
│   ├── setting/             # Configuration management
│   ├── logger/              # Logging utilities
│   └── utils/               # Utility helper functions
├── configs/                 # Configuration files
├── migrations/              # Database migration scripts (SQL)
├── scripts/                 # Build and utility scripts
├── tests/                   # Test files & test utilities
├── docs/                    # Project documentation
├── third_party/             # External dependencies/libraries
├── response/                # HTTP response utilities
├── global/                  # Global constants and application-wide variables
├── CLAUDE.md                # Claude Code context and guidance
├── SETUP.md                 # Detailed setup guide
└── go.mod                   # Go module definition
```

## 🚀 Quick Start

### Prerequisites
- Go 1.26+ installed
- (Database setup optional - structure ready in `internal/repo/`)

### Running the Server

```bash
go run cmd/server/main.go
```

Server starts on `http://localhost:8080`

**Test endpoints:**
```bash
curl http://localhost:8080/v1/ping
curl http://localhost:8080/v1/query?name=John
curl http://localhost:8080/v1/user/alice
curl http://localhost:8080/v2/ping
curl http://localhost:8080/v2/default-query
```

### Running Background Jobs
```bash
go run cmd/cronjob/main.go
```

### Running CLI Tools
```bash
go run cmd/cli/main.go [command] [options]
```

## 🏗️ Architecture

The project follows **Clean Architecture** with clear separation of concerns:

```
HTTP Request → Router → Controller → Service → Repository → Database
```

**Layers:**
- **Router** (`internal/routers/`) - Route definitions and request dispatch
- **Controller** (`internal/controller/`) - HTTP request/response handling
- **Service** (`internal/service/`) - Business logic and core functionality
- **Repository** (`internal/repo/`) - Data access layer (DAO pattern)
- **Models** (`internal/models/`) - Domain models and DTOs

**Middlewares** (`internal/middlewares/`) handle cross-cutting concerns like authentication and logging.

## 🌐 Web Framework: Gin Gonic

Built with [Gin Gonic](https://github.com/gin-gonic/gin), a fast and lightweight HTTP framework.

**Current API Implementation:**
- **v1 endpoints:** `/v1/ping`, `/v1/user/:name`, `/v1/query?name=...`
- **v2 endpoints:** `/v2/ping`, `/v2/default-query?name=...` (with default value)

Routes are defined in `internal/routers/router.go` with parameter handling for:
- Query parameters: `c.Query("param")`
- Path parameters: `c.Param("param")`
- Default values: `c.DefaultQuery("param", "default")`
- JSON bodies: `c.BindJSON(&struct)`
- Headers: `c.GetHeader("X-Header")`

## 📝 Development Guide

**To add a new endpoint:**

1. Define the handler in `internal/routers/router.go` (or move to `internal/controller/` as it grows)
2. Register the route in the appropriate API version group
3. Extract parameters using Gin context methods
4. Return responses with `c.JSON(statusCode, data)`

**Example:**
```go
v1.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.JSON(http.StatusOK, gin.H{"user_name": name})
})
```

## 📦 Configuration

Configuration files are in the `configs/` directory. The `pkg/setting` package handles configuration management.

## 🔄 Database Migrations

Database migration scripts are in `migrations/`. Use `golang-migrate` or similar tools:

```bash
migrate -path migrations -database "your-db-url" up
```

## 🧪 Testing

```bash
go test ./...              # Run all tests
go test -v ./...           # Verbose output
go test -cover ./...       # Show code coverage
```

Tests can be added to the `tests/` directory.

## 📚 Documentation

- **CLAUDE.md** - Development context for Claude Code
- **SETUP.md** - Detailed directory structure and setup guide
- **docs/** - Additional project documentation

## 📄 License

Add your license information here.
