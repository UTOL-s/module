# Dynamic GORM Module

A highly configurable and dynamic GORM module for the UTOL module system that supports multiple database types, connection pooling, logging, and more.

## Features

- **Multi-Database Support**: PostgreSQL, MySQL, SQLite, and SQL Server
- **Dynamic Configuration**: Runtime configuration through YAML files and environment variables
- **Connection Pooling**: Configurable connection pool settings
- **Advanced Logging**: Customizable logging levels and slow query detection
- **Debug Mode**: Dry run mode for development and testing
- **Dependency Injection**: Seamless integration with Uber FX

## Supported Database Types

### PostgreSQL
```yaml
database:
  type: "postgres"
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "utol_db"
  sslmode: "disable"
```

### MySQL
```yaml
database:
  type: "mysql"
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  dbname: "utol_db"
  charset: "utf8mb4"
  parse_time: true
  loc: "Local"
```

### SQLite
```yaml
database:
  type: "sqlite"
  file: "./data/utol.db"
```

### SQL Server
```yaml
database:
  type: "sqlserver"
  host: "localhost"
  port: 1433
  user: "sa"
  password: "password"
  dbname: "utol_db"
```

## Configuration Options

### Connection Pool Settings
```yaml
database:
  pool:
    max_idle_conns: 10        # Maximum number of idle connections
    max_open_conns: 100       # Maximum number of open connections
    conn_max_lifetime: 3600   # Maximum lifetime of connections (seconds)
    conn_max_idle_time: 600   # Maximum idle time of connections (seconds)
```

### Logging Configuration
```yaml
database:
  log:
    level: 4                                    # Log level (1=Silent, 2=Error, 3=Warn, 4=Info)
    slow_threshold: 5000                       # Slow query threshold (milliseconds)
    colorful: true                             # Enable colored output
    ignore_record_not_found_error: true        # Ignore "record not found" errors
```

### Debug Mode
```yaml
database:
  debug: false  # Enable dry run mode (no actual database operations)
```

## Usage

### Basic Usage
```go
package main

import (
    "github.com/UTOL-s/module/fxGorm"
    "go.uber.org/fx"
)

func main() {
    app := fx.New(
        fxGorm.FxGorm,
        // ... other modules
    )
    
    app.Run()
}
```

### Advanced Usage with Custom Configuration
```go
package main

import (
    "github.com/UTOL-s/module/fxGorm"
    "go.uber.org/fx"
)

type App struct {
    DB *gorm.DB
}

func NewApp(db *gorm.DB) *App {
    return &App{DB: db}
}

func main() {
    app := fx.New(
        fxGorm.FxGorm,
        fx.Provide(NewApp),
        fx.Invoke(func(app *App) {
            // Your application logic here
        }),
    )
    
    app.Run()
}
```

## Environment Variables

All configuration options can be overridden using environment variables:

```bash
# Database type
export DATABASE_TYPE=postgres

# Connection settings
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASSWORD=password
export DATABASE_DBNAME=utol_db

# Pool settings
export DATABASE_POOL_MAX_IDLE_CONNS=10
export DATABASE_POOL_MAX_OPEN_CONNS=100

# Logging settings
export DATABASE_LOG_LEVEL=4
export DATABASE_LOG_SLOW_THRESHOLD=5000
```

## Default Values

If configuration options are not specified, the following defaults are used:

- **Database Type**: PostgreSQL (if not specified)
- **Max Idle Connections**: 10
- **Max Open Connections**: 100
- **Connection Max Lifetime**: 1 hour
- **Connection Max Idle Time**: 10 minutes
- **Log Level**: Info (4)
- **Slow Threshold**: 5 seconds
- **Debug Mode**: false

## Error Handling

The module includes comprehensive error handling:

- Database connection failures
- Configuration validation
- Connection pool setup errors
- Database ping failures

All errors are wrapped with context for better debugging.

## Testing

The module includes built-in connection testing:

- Automatic connection validation on startup
- Database ping verification
- Configuration validation

## Migration from Previous Version

If you're upgrading from the previous version, update your configuration:

```yaml
# Old configuration
database:
  host: "localhost"
  port: 5432
  # ... other settings

# New configuration
database:
  type: "postgres"  # Add this line
  host: "localhost"
  port: 5432
  # ... other settings
```

## Dependencies

The module requires the following GORM drivers (automatically added to go.mod):

- `gorm.io/driver/postgres` - PostgreSQL support
- `gorm.io/driver/mysql` - MySQL support
- `gorm.io/driver/sqlite` - SQLite support
- `gorm.io/driver/sqlserver` - SQL Server support

## License

This module is part of the UTOL project and follows the same license terms. 