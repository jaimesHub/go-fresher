# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Running the Application

**Start the HTTP API server:**
```bash
go run cmd/server/main.go
```
Server listens on `http://localhost:8080` by default. Test with: `curl http://localhost:8080/v1/ping`

**Run background jobs:**
```bash
go run cmd/cronjob/main.go
```

**Run CLI tools:**
```bash
go run cmd/cli/main.go [command] [options]
```

### Testing

Currently no test files exist in the repository. To add tests:
```bash
go test ./...                    # Run all tests
go test -v ./...                 # Verbose test output
go test -run TestName ./...      # Run a specific test
go test -cover ./...             # Show code coverage
```

### Build & Dependencies

**Install or update dependencies:**
```bash
go mod tidy                      # Clean up unused dependencies
go get -u github.com/...         # Update a package
```

**Build binary:**
```bash
go build -o fresher cmd/server/main.go  # Build server binary
```

**Run formatter and linter:**
```bash
go fmt ./...                     # Format all Go files
gofmt -w .                       # Alternative format command
```

## Architecture

This project follows **Clean Architecture** principles with the following layers:

### Directory Structure

```
internal/              # Private application code (not exported)
├── routers/          # HTTP route definitions → where new API endpoints are registered
├── controller/       # HTTP request handlers (empty - add handlers here)
├── service/          # Business logic layer (empty - core functionality goes here)
├── models/           # Domain models and data structures (empty)
├── repo/             # Data access layer / DAO pattern (empty)
├── middlewares/      # HTTP middleware - auth, logging, error handling (empty)
└── initialize/       # Application startup and initialization (empty)

pkg/                   # Reusable utilities (can be shared across projects)
├── setting/          # Configuration management
├── logger/           # Logging utilities
└── utils/            # Utility helper functions

cmd/                   # Executable entry points
├── server/           # HTTP API server (main.go)
├── cronjob/          # Background job executor
└── cli/              # CLI tools

configs/              # Configuration files
migrations/           # Database migration scripts (SQL)
response/             # HTTP response utilities
global/               # Global constants and variables
tests/                # Test files and test utilities
```

### Dependency Inversion Flow

When adding a new feature:

1. **Router** (`internal/routers/router.go`): Register new routes/endpoints
   - Example: `r.GET("/api/users", controllers.GetUsers)`

2. **Controller** (`internal/controller/`): Handle HTTP request/response
   - Extract parameters, validate input
   - Call service layer for business logic
   - Return responses (use `c.JSON()` in Gin)

3. **Service** (`internal/service/`): Implement business logic
   - Contains the core application logic
   - Calls repository for data access

4. **Repository** (`internal/repo/`): Data access layer
   - Database queries, caching logic
   - Abstracts data source from service layer

5. **Models** (`internal/models/`): Data structures
   - Domain models, request/response DTOs

### Web Framework: Gin

The project uses **Gin Gonic** (`github.com/gin-gonic/gin`), a fast HTTP web framework.

**Key Gin patterns used in this project:**

```go
// Create router
r := gin.Default()  // Includes logger and recovery middleware

// Route groups for API versioning
v1 := r.Group("/v1")
v1.GET("/ping", handler)

// Parameter handling in handlers
func handler(c *gin.Context) {
    // Query parameters: /endpoint?name=value
    name := c.Query("name")
    
    // Path parameters: /user/:name
    name := c.Param("name")
    
    // Default query value
    name := c.DefaultQuery("name", "Guest")
    
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
```

Gin routes are defined in `internal/routers/router.go` and currently handles `/v1/*` and `/v2/*` endpoints.

## Adding New Endpoints

To add a new endpoint:

1. **Create handler** in `internal/controller/` (or inline in `internal/routers/router.go` for now)
2. **Register route** in `internal/routers/router.go` with appropriate HTTP method and path
3. **Extract parameters** using Gin context methods: `c.Query()`, `c.Param()`, `c.BindJSON()`
4. **Return response** using `c.JSON(statusCode, data)`

Example from the codebase:
```go
// In internal/routers/router.go
v1.GET("/user/:name", UserHandler)

func UserHandler(c *gin.Context) {
    name := c.Param("name")
    c.JSON(http.StatusOK, gin.H{"user_name": name})
}
```

## Project Status

- **Web framework**: Gin (initialized, basic routing working)
- **Database integration**: Not yet implemented (structure exists at `internal/repo/`)
- **Service layer**: Not yet implemented (structure exists at `internal/service/`)
- **Tests**: Not yet implemented (use `go test` when added)
- **Middleware**: Not yet implemented (structure exists at `internal/middlewares/`)
