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

## Releasing

This project uses GitHub Actions for automated releases with separate workflows for each module.

### Release Types

#### 1. Complete Module Release
Releases the entire module with all fx modules included.

**Tag-based Release:**
```bash
git tag v1.0.0
git push origin v1.0.0
```

**Manual Release:**
1. Go to GitHub Actions → "Release Complete Module"
2. Click "Run workflow"
3. Enter version (e.g., `v1.0.0`)

#### 2. fxConfig Module Release
Releases only the fxConfig module independently.

**Tag-based Release:**
```bash
git tag fxconfig-v1.0.0
git push origin fxconfig-v1.0.0
```

**Manual Release:**
1. Go to GitHub Actions → "Release fxConfig Module"
2. Click "Run workflow"
3. Enter version (e.g., `fxconfig-v1.0.0`)

#### 3. fxEcho Module Release
Releases only the fxEcho module independently.

**Tag-based Release:**
```bash
git tag fxecho-v1.0.0
git push origin fxecho-v1.0.0
```

**Manual Release:**
1. Go to GitHub Actions → "Release fxEcho Module"
2. Click "Run workflow"
3. Enter version (e.g., `fxecho-v1.0.0`)

### Workflow Features

Each release workflow:
1. ✅ Runs tests for the specific module
2. 🔍 Performs code quality checks (vet, formatting)
3. 🏗️ Builds the module
4. 🏷️ Creates a GitHub release with changelog
5. 📦 Publishes to Go module proxy

### Installation by Version

```bash
# Complete module
go get github.com/UTOL-s/module@v1.0.0

# Individual modules
go get github.com/UTOL-s/module/fxConfig@fxconfig-v1.0.0
go get github.com/UTOL-s/module/fxEcho@fxecho-v1.0.0
```

## License

This project is licensed under the terms specified in the LICENSE file. 