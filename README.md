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

## Continuous Integration

### Automated Testing
All GitHub Actions workflows run on:
- **All branches** - For continuous testing and validation
- **Pull requests** to main/master - For code review validation
- **Tags** - For releases

### Workflow Jobs
Each workflow includes:
1. **Test** - Runs tests, vet, and formatting checks
2. **Build** - Builds the module and verifies it
3. **Release** - Creates GitHub releases (only on tags or manual dispatch)

## Releasing

This project uses automated releases based on [Conventional Commits](https://www.conventionalcommits.org/) and [semantic-release](https://github.com/semantic-release/semantic-release), similar to the [UTOL-s/stoken](https://github.com/UTOL-s/stoken) repository.

### Commit Message Conventions

Use these prefixes in your commit messages to trigger automatic releases:

- **`feat:`** - New features (triggers minor version bump)
- **`fix:`** - Bug fixes (triggers patch version bump)
- **`perf:`** - Performance improvements (triggers patch version bump)
- **`BREAKING CHANGE:`** - Breaking changes (triggers major version bump)

#### Examples:
```bash
git commit -m "feat: add new config loader"
git commit -m "fix: correct environment variable parsing"
git commit -m "perf: optimize database connection pooling"
git commit -m "feat: add new middleware

BREAKING CHANGE: config structure has changed"
```

### Release Process

#### Automatic Releases
When you push commits to the `main` branch with conventional commit messages:

1. **Tests run** - All tests and code quality checks are executed
2. **Version analysis** - semantic-release analyzes commit messages
3. **Version bump** - Determines the next semantic version
4. **Tag creation** - Creates a new tag with the appropriate format
5. **GitHub release** - Creates a GitHub release with changelog
6. **Go module proxy** - Publishes to the Go module proxy

#### Release Types

##### fxConfig Module
- **Tag format**: `fxconfig-v1.2.3`
- **Trigger**: Commits affecting `./fxConfig` directory
- **Installation**: `go get github.com/UTOL-s/module/fxConfig@fxconfig-v1.2.3`

##### fxEcho Module
- **Tag format**: `fxecho-v1.2.3`
- **Trigger**: Commits affecting `./fxEcho` directory
- **Installation**: `go get github.com/UTOL-s/module/fxEcho@fxecho-v1.2.3`

##### Complete Module
- **Tag format**: `v1.2.3`
- **Trigger**: Any commits to the repository
- **Installation**: `go get github.com/UTOL-s/module@v1.2.3`

### Manual Releases (if needed)

If you need to manually release a version:

```bash
# Create and push a tag
git tag fxconfig-v1.0.0
git push origin fxconfig-v1.0.0
```

### Version Bumping Rules

- **Major version** (`1.0.0` → `2.0.0`): Breaking changes
- **Minor version** (`1.0.0` → `1.1.0`): New features
- **Patch version** (`1.0.0` → `1.0.1`): Bug fixes and performance improvements

### Installation by Version

```bash
# Complete module
go get github.com/UTOL-s/module@v1.0.0

# Individual modules
go get github.com/UTOL-s/module/fxConfig@fxconfig-v1.0.0
go get github.com/UTOL-s/module/fxEcho@fxecho-v1.0.0

# Latest versions
go get github.com/UTOL-s/module@latest
go get github.com/UTOL-s/module/fxConfig@latest
go get github.com/UTOL-s/module/fxEcho@latest
```

## License

This project is licensed under the terms specified in the LICENSE file. 