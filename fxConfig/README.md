# FX Config Module

A flexible configuration management module for the UTOL module system that provides YAML-based configuration with environment variable support and dependency injection integration.

## Features

- **YAML Configuration**: Easy-to-read YAML configuration files
- **Environment Variable Support**: Override any config value with environment variables
- **Dependency Injection**: Seamless integration with Uber FX
- **Type-Safe Access**: Strongly typed configuration access methods
- **Automatic Environment Expansion**: Support for environment variable expansion in config files
- **Multiple Access Patterns**: Both struct-based and accessor-based configuration access

## Configuration Structure

The module supports a hierarchical configuration structure:

```yaml
app:
  name: "UTOL Application"
  port: "8080"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "utol_db"
  sslmode: "disable"
```

## Usage

### Basic Usage with FX

```go
package main

import (
    "github.com/UTOL-s/module/fxConfig"
    "go.uber.org/fx"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fx.Invoke(func(config *fxconfig.Config) {
            // Access configuration
            appName := config.App.Name
            dbHost := config.Database.Host
        }),
    )
    
    app.Run()
}
```

### Using Config Accessor (Yokai-style)

```go
package main

import (
    "github.com/UTOL-s/module/fxConfig"
    "go.uber.org/fx"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fx.Invoke(func(accessor *fxconfig.Accessor) {
            // Access configuration using dot notation
            appName := accessor.String("app.name")
            dbPort := accessor.Int("database.port")
            debugMode := accessor.Bool("app.debug")
        }),
    )
    
    app.Run()
}
```

### Advanced Usage with Custom Structs

```go
package main

import (
    "github.com/UTOL-s/module/fxConfig"
    "go.uber.org/fx"
)

type AppConfig struct {
    Config *fxconfig.Config
}

func NewAppConfig(config *fxconfig.Config) *AppConfig {
    return &AppConfig{Config: config}
}

func (ac *AppConfig) GetDatabaseDSN() string {
    return ac.Config.PostgresDSN()
}

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fx.Provide(NewAppConfig),
        fx.Invoke(func(appConfig *AppConfig) {
            dsn := appConfig.GetDatabaseDSN()
            // Use DSN for database connection
        }),
    )
    
    app.Run()
}
```

## Configuration Access Methods

### Struct-Based Access

```go
config := &fxconfig.Config{}

// Access app configuration
appName := config.App.Name
appPort := config.App.Port

// Access database configuration
dbHost := config.Database.Host
dbPort := config.Database.Port
dbUser := config.Database.User
dbPassword := config.Database.Password
dbName := config.Database.DBName
dbSSLMode := config.Database.SSLMode
```

### Accessor-Based Access

```go
accessor := &fxconfig.Accessor{}

// String values
appName := accessor.String("app.name")
dbHost := accessor.String("database.host")

// Integer values
dbPort := accessor.Int("database.port")

// Boolean values
debugMode := accessor.Bool("app.debug")

// Float values
timeout := accessor.Float64("app.timeout")

// All settings
allSettings := accessor.AllSettings()
```

## Environment Variable Support

All configuration values can be overridden using environment variables. The module automatically converts dot notation to underscore notation:

```bash
# Override app configuration
export APP_NAME="My Custom App"
export APP_PORT="9090"

# Override database configuration
export DATABASE_HOST="prod-db.example.com"
export DATABASE_PORT="5432"
export DATABASE_USER="prod_user"
export DATABASE_PASSWORD="prod_password"
export DATABASE_DBNAME="prod_db"
export DATABASE_SSLMODE="require"
```

## Environment Variable Expansion

The module supports environment variable expansion within the YAML configuration file:

```yaml
app:
  name: "${APP_NAME:-UTOL Application}"
  port: "${APP_PORT:-8080}"

database:
  host: "${DATABASE_HOST:-localhost}"
  port: ${DATABASE_PORT:-5432}
  user: "${DATABASE_USER:-postgres}"
  password: "${DATABASE_PASSWORD:-password}"
  dbname: "${DATABASE_DBNAME:-utol_db}"
  sslmode: "${DATABASE_SSLMODE:-disable}"
```

## Database DSN Generation

The module provides a convenient method to generate PostgreSQL DSN strings:

```go
config := &fxconfig.Config{
    Database: struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        DBName   string `mapstructure:"dbname"`
        SSLMode  string `mapstructure:"sslmode"`
    }{
        Host:     "localhost",
        Port:     5432,
        User:     "postgres",
        Password: "password",
        DBName:   "utol_db",
        SSLMode:  "disable",
    },
}

dsn := config.PostgresDSN()
// Result: "host=localhost user=postgres password=password dbname=utol_db port=5432 sslmode=disable"
```

## File Structure

The module expects the following file structure:

```
your-app/
├── configs/
│   └── config.yaml
├── main.go
└── ...
```

## Error Handling

The module provides comprehensive error handling:

- Configuration file not found
- Invalid YAML syntax
- Environment variable expansion failures
- Configuration unmarshaling errors

All errors are wrapped with context for better debugging.

## Testing

The module includes built-in testing support:

```go
package main

import (
    "testing"
    "github.com/UTOL-s/module/fxConfig"
)

func TestConfig(t *testing.T) {
    config, err := fxconfig.NewConfig()
    if err != nil {
        t.Fatalf("Failed to load config: %v", err)
    }
    
    if config.App.Name == "" {
        t.Error("App name should not be empty")
    }
}
```

## Dependencies

- `github.com/spf13/viper`: Configuration management
- `go.uber.org/fx`: Dependency injection framework

## Migration Guide

If you're migrating from a different configuration system:

1. Create a `configs/config.yaml` file with your configuration
2. Replace direct environment variable access with the config accessor
3. Update your dependency injection setup to include `fxConfig.FxConfig`
4. Replace hardcoded values with configuration references

## Best Practices

1. **Use Environment Variables for Secrets**: Never commit passwords or API keys to version control
2. **Provide Sensible Defaults**: Use environment variable expansion with defaults
3. **Validate Configuration**: Add validation logic for critical configuration values
4. **Use Type-Safe Access**: Prefer struct-based access for known configuration keys
5. **Centralize Configuration**: Use the accessor pattern for dynamic configuration access
