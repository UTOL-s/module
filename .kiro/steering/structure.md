# Project Structure & Organization

## Repository Layout

### Root Level Structure
```
module/
├── fxConfig/           # Configuration management module
├── fxEcho/             # HTTP server module with Echo integration
├── fxGorm/             # Database ORM module with GORM
├── fxSupertoken/       # Authentication module with SuperTokens
├── api_rest/           # Complete integrated REST API example
├── configs/            # Global configuration templates
├── internal/           # Shared internal packages
├── .air.toml           # Air hot reload configuration
├── Makefile            # Root build automation
├── run-api.sh          # Development runner script
├── go.mod              # Root Go module definition
└── README.md           # Main project documentation
```

## Module Organization Pattern

Each module follows a consistent structure:

### Standard Module Layout
```
fxModuleName/
├── README.md           # Module-specific documentation
├── module.go           # FX module definition and providers
├── config.go           # Configuration structures (if applicable)
├── types.go            # Type definitions and interfaces
├── example_test.go     # Usage examples and integration tests
├── module_test.go      # Unit tests for the module
└── example/            # Standalone example application
    ├── main.go
    ├── go.mod
    └── configs/
```

## API REST Example Structure

The `api_rest` directory demonstrates complete application architecture:

```
api_rest/
├── cmd/                # CLI commands using Cobra
│   ├── root.go         # Root command definition
│   └── serve.go        # Server start command
├── configs/            # Configuration files
│   └── config.yaml     # Application configuration
├── internal/           # Private application code
│   ├── bootstrap.go    # Application initialization
│   ├── register.go     # Dependency registration
│   ├── router.go       # Route definitions
│   ├── handler/        # HTTP request handlers
│   │   ├── health.go   # Health check endpoints
│   │   └── user.go     # User management endpoints
│   ├── model/          # Data models and structures
│   │   └── user.go     # User model definition
│   ├── repository/     # Data access layer
│   │   └── user_repository.go
│   └── service/        # Business logic layer
│       └── user_service.go
├── tmp/                # Temporary build artifacts (Air)
├── main.go             # Application entry point
├── go.mod              # Module dependencies
├── Makefile            # Build automation
├── docker-compose.yaml # Docker development setup
└── .env                # Environment variables
```

## Architecture Patterns

### Dependency Injection Pattern
- All modules use Uber FX for dependency injection
- Providers are defined in `module.go` files
- Dependencies are injected through constructor functions
- Lifecycle management is handled by FX

### Clean Architecture Layers
```
Handler Layer    → HTTP request/response handling
Service Layer    → Business logic and validation
Repository Layer → Data access and persistence
Model Layer      → Data structures and entities
```

### Configuration Management
- YAML-based configuration with environment variable overrides
- Configuration structures use `mapstructure` tags
- Environment variables follow `SECTION_KEY` naming convention
- Default values provided through struct tags or code

## File Naming Conventions

### Go Files
- `module.go`: FX module definition and main providers
- `config.go`: Configuration structures and validation
- `types.go`: Interface definitions and type declarations
- `*_test.go`: Test files following Go conventions
- `example_test.go`: Integration tests and usage examples

### Directories
- `internal/`: Private packages not intended for external use
- `cmd/`: Command-line interface implementations
- `configs/`: Configuration files and templates
- `example/`: Standalone example applications
- `tmp/`: Temporary files and build artifacts (gitignored)

## Import Organization

### Import Grouping (Standard Go Practice)
```go
import (
    // Standard library
    "context"
    "fmt"
    "net/http"
    
    // Third-party packages
    "github.com/labstack/echo/v4"
    "go.uber.org/fx"
    "go.uber.org/zap"
    
    // Local packages
    "github.com/UTOL-s/module/fxConfig"
    "github.com/UTOL-s/module/fxEcho"
)
```

## Testing Structure

### Test Organization
- Unit tests alongside source files (`*_test.go`)
- Integration tests in `example_test.go` files
- Test utilities and helpers in dedicated test packages
- Mock implementations for external dependencies

### Test Naming
- `TestFunctionName` for unit tests
- `ExampleFunctionName` for example tests
- `BenchmarkFunctionName` for performance tests

## Configuration File Locations

### Development Configuration
- `configs/config.yaml.example`: Template configuration
- `api_rest/configs/config.yaml`: API example configuration
- `module/example/configs/config.yaml`: Module-specific examples

### Environment Files
- `.env`: Local development environment variables
- `env.example`: Template for environment variables
- Docker Compose environment configuration

## Build Artifacts

### Generated Files (Gitignored)
- `tmp/`: Air temporary files and build artifacts
- `*.log`: Log files from development tools
- Binary executables (e.g., `api_rest_optimized`)
- `.DS_Store`: macOS system files

### Version Control
- All source code and configuration templates are tracked
- Build artifacts and temporary files are gitignored
- Environment files with secrets are gitignored
- Documentation and examples are tracked

## Module Interdependencies

### Dependency Flow
```
api_rest → fxConfig + fxEcho + fxGorm + fxSupertoken
fxEcho → fxConfig (optional)
fxGorm → fxConfig (optional)
fxSupertoken → fxConfig (optional)
```

### Module Isolation
- Each module can be used independently
- Minimal cross-module dependencies
- Clear interfaces between modules
- Optional configuration integration