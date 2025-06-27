# Unified Transport Operations League - Go Module

This Go module provides fx modules for the Unified Transport Operations League project.

## Modules

### fxConfig
A configuration module using Uber's fx dependency injection framework with Viper for configuration management.

### fxEcho
An Echo web framework module for Uber's fx dependency injection.

## Installation

### Complete Module
```bash
go get github.com/UTOL-s/module
```

### Individual Modules
```bash
# fxConfig only
go get github.com/UTOL-s/module/fxConfig

# fxEcho only
go get github.com/UTOL-s/module/fxEcho
```

## Usage

### Complete Module
```go
package main

import (
    "go.uber.org/fx"
    "github.com/UTOL-s/module/fxConfig"
    "github.com/UTOL-s/module/fxEcho"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fxEcho.FxEcho,
        // Add your application components here
    )
    
    app.Run()
}
```

### Individual Modules
```go
// Using only fxConfig
import (
    "go.uber.org/fx"
    "github.com/UTOL-s/module/fxConfig"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        // ... other modules
    )
    app.Run()
}
```

## Configuration

The fxConfig module expects a `config.yaml` file in the `./configs/` directory:

```yaml
app:
  name: "your-app-name"
  port: "8080"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "database"
  sslmode: "disable"
```

Environment variables can be used to override configuration values using the pattern `APP_NAME`, `DATABASE_HOST`, etc.

## Development

### Prerequisites
- Go 1.24.2 or later

### Running Tests
```bash
# Test all modules
go test -v ./...

# Test specific modules
go test -v ./fxConfig
go test -v ./fxEcho
```

### Building
```bash
# Build all modules
go build -v ./...

# Build specific modules
go build -v ./fxConfig
go build -v ./fxEcho
```

## How to Push Updates to This Go Module

### CI/CD Process

This repository uses GitHub Actions for continuous integration and delivery, similar to the [UTOL-s/stoken](https://github.com/UTOL-s/stoken) repository:

#### Pull Request Workflow

When a pull request is opened against the main branch, the following checks are automatically run:

* Dependency verification
* Unit tests for all fx modules
* Code formatting checks
* Static analysis with go vet
* go mod tidy verification

#### Release Workflow

When changes are merged to the main branch, a release workflow is triggered that:

1. Runs tests to ensure code quality
2. Determines the next semantic version based on commit history
3. Creates a new tag and GitHub release
4. Publishes the module to the Go package registry

### Add Prefixes on commit

* feature
* feat
* fix
* bugfix
* perf
* refactor
* test
* breaking
* major

```
Example:
   feature: additional login
   feat: add new config loader
   fix: correct environment variable parsing
   perf: optimize database connection pooling
   breaking: change config structure
```

### Manual Release Process (if needed)

If you need to manually release a new version:

1. **Clone the repository** (if you haven't already):  
   git clone https://github.com/UTOL-s/module.git  
   cd module
2. **Make your changes** to the code.
3. **Update dependencies** if necessary:  
   go mod tidy
4. **Commit your changes**:  
   git add .  
   git commit -m "Description of your changes"
5. **Push your changes** to trigger the automated release:  
   git push origin main
6. **Verify the new version** after the GitHub Action completes:  
   go list -m github.com/UTOL-s/module@<new-version>

### Best Practices for Go Module Versioning

* Follow Semantic Versioning (SemVer) for your tags.
* Major version changes (v1 → v2) that include breaking changes should use a different module path (e.g., `/v2` suffix).
* Include a CHANGELOG.md to document changes between versions.
* Use go.mod's `replace` directive during local development if needed.

## Project Structure

The module is organized with the following structure:

```
module/
├── .github/                  # GitHub specific files
│   └── workflows/            # CI/CD workflow definitions
│       ├── release.yml       # Release workflow
│       └── test.yml          # Test workflow
├── configs/                  # Configuration files
│   └── config.yaml.example   # Example configuration
├── fxConfig/                 # Configuration module
│   ├── config.go            # Configuration implementation
│   ├── config_test.go       # Configuration tests
│   └── module.go            # fx module definition
├── fxEcho/                   # Echo module
│   ├── module.go            # fx module definition
│   └── module_test.go       # Echo module tests
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
└── README.md                 # Documentation
```

## License

This project is licensed under the terms specified in the LICENSE file. 