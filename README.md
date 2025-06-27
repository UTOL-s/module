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

This repository uses GitHub Actions for continuous integration and delivery, similar to the [UTOL-s/stoken](https://github.com/UTOL-s/stoken) repository. Each fx module has its own separate workflows for independent testing and releases.

#### GitHub Actions Workflows

##### Module-Level Workflows
- **`test.yml`** - Tests the overall module structure (excludes fx-specific changes)
- **`release.yml`** - Releases the complete module (excludes fx-specific changes)

##### fxConfig Module Workflows
- **`fxconfig-test.yml`** - Tests fxConfig module specifically
- **`fxconfig-release.yml`** - Releases fxConfig module independently

##### fxEcho Module Workflows
- **`fxecho-test.yml`** - Tests fxEcho module specifically
- **`fxecho-release.yml`** - Releases fxEcho module independently

#### Workflow Triggers

- **fxConfig changes** → Triggers `fxconfig-test.yml` and `fxconfig-release.yml`
- **fxEcho changes** → Triggers `fxecho-test.yml` and `fxecho-release.yml`
- **Module-wide changes** → Triggers `test.yml` and `release.yml`

#### Pull Request Workflow

When a pull request is opened against the main branch, the following checks are automatically run based on what files were changed:

* Dependency verification
* Unit tests for affected modules
* Code formatting checks
* Static analysis with go vet
* go mod tidy verification

#### Release Workflow

When changes are merged to the main branch, release workflows are triggered based on what was changed:

1. Runs tests to ensure code quality
2. Uses `ietf-tools/semver-action@v1` to determine the next semantic version based on conventional commits
3. Creates a new tag and GitHub release
4. Publishes the module to the Go package registry

### Semantic Versioning with ietf-tools/semver-action@v1

This project uses the `ietf-tools/semver-action@v1` for robust semantic versioning based on [Conventional Commits](https://www.conventionalcommits.org/). The action automatically analyzes commit messages to determine the appropriate version bump.

#### Supported Commit Types

The action supports the following conventional commit types:

- **`feat:`** - New features (triggers minor version bump)
- **`fix:`** - Bug fixes (triggers patch version bump)
- **`docs:`** - Documentation changes (triggers patch version bump)
- **`style:`** - Code style changes (triggers patch version bump)
- **`refactor:`** - Code refactoring (triggers patch version bump)
- **`perf:`** - Performance improvements (triggers patch version bump)
- **`test:`** - Adding or updating tests (triggers patch version bump)
- **`chore:`** - Maintenance tasks (triggers patch version bump)
- **`ci:`** - CI/CD changes (triggers patch version bump)
- **`build:`** - Build system changes (triggers patch version bump)
- **`BREAKING CHANGE:`** - Breaking changes (triggers major version bump)

#### Commit Message Examples

```
feat: add new configuration loader
fix: correct environment variable parsing
docs: update README with new examples
style: format code according to standards
refactor: restructure configuration loading
perf: optimize database connection pooling
test: add unit tests for config validation
chore: update dependencies
ci: add new GitHub Actions workflow
build: update Go version requirement
feat: add new middleware

BREAKING CHANGE: config structure has changed
```

### Versioning Strategy

Each module uses its own versioning scheme with the `ietf-tools/semver-action@v1`:

- **Complete Module**: `v1.2.3` (e.g., `v1.0.0`)
- **fxConfig Module**: `fxconfig-v1.2.3` (e.g., `fxconfig-v1.0.0`)
- **fxEcho Module**: `fxecho-v1.2.3` (e.g., `fxecho-v1.0.0`)

### Manual Release Process (if needed)

If you need to manually release a new version:

1. **Clone the repository** (if you haven't already):  
   git clone https://github.com/UTOL-s/module.git  
   cd module
2. **Make your changes** to the code.
3. **Update dependencies** if necessary:  
   go mod tidy
4. **Commit your changes** using conventional commit format:  
   git add .  
   git commit -m "feat: add new feature"
5. **Push your changes** to trigger the automated release:  
   git push origin main
6. **Verify the new version** after the GitHub Action completes:  
   go list -m github.com/UTOL-s/module@<new-version>

### Best Practices for Go Module Versioning

* Follow Semantic Versioning (SemVer) for your tags.
* Use conventional commit messages to enable automatic versioning.
* Major version changes (v1 → v2) that include breaking changes should use a different module path (e.g., `/v2` suffix).
* Include a CHANGELOG.md to document changes between versions.
* Use go.mod's `replace` directive during local development if needed.

## Project Structure

The module is organized with the following structure:

```
module/
├── .github/                  # GitHub specific files
│   └── workflows/            # CI/CD workflow definitions
│       ├── test.yml          # Module-level test workflow
│       ├── release.yml       # Module-level release workflow
│       ├── fxconfig-test.yml # fxConfig test workflow
│       ├── fxconfig-release.yml # fxConfig release workflow
│       ├── fxecho-test.yml   # fxEcho test workflow
│       └── fxecho-release.yml # fxEcho release workflow
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
├── CHANGELOG.md              # Version change log
└── README.md                 # Documentation
```

## License

This project is licensed under the terms specified in the LICENSE file. 