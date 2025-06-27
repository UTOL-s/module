# FX Echo Module

A module for integrating the Echo web framework with Uber FX dependency injection in the UTOL module system.

## Overview

The fxEcho module is designed to provide seamless integration between the Echo web framework and the FX dependency injection system. This module will handle Echo server setup, middleware configuration, and route management through dependency injection.

## Current Status

⚠️ **This module is currently under development and contains placeholder functionality.**

The module is set up with the basic structure for FX integration but requires implementation of Echo-specific functionality.

## Planned Features

- **Echo Server Setup**: Automatic Echo server initialization with FX
- **Middleware Integration**: Dependency injection for Echo middleware
- **Route Management**: FX-based route registration and management
- **Configuration Integration**: Integration with fxConfig for server settings
- **Graceful Shutdown**: Proper server shutdown handling with FX lifecycle
- **Health Checks**: Built-in health check endpoints
- **CORS Support**: Configurable CORS middleware
- **Request Logging**: Structured request logging middleware

## Expected Usage

```go
package main

import (
    "github.com/UTOL-s/module/fxEcho"
    "github.com/UTOL-s/module/fxConfig"
    "go.uber.org/fx"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fxEcho.FxEcho,
        fx.Invoke(func(e *echo.Echo) {
            // Echo server will be automatically started
        }),
    )
    
    app.Run()
}
```

## Expected Configuration

The module will likely support configuration through the fxConfig module:

```yaml
server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 60s
  
middleware:
  cors:
    enabled: true
    origins: ["*"]
  logger:
    enabled: true
    level: "info"
  recovery:
    enabled: true
```

## Dependencies

- `github.com/labstack/echo/v4`: Echo web framework
- `go.uber.org/fx`: Dependency injection framework

## Development Roadmap

1. **Phase 1**: Basic Echo server setup with FX
2. **Phase 2**: Middleware integration
3. **Phase 3**: Route management
4. **Phase 4**: Advanced features (health checks, metrics, etc.)

## Contributing

This module is open for contributions. Please check the main project repository for contribution guidelines.

## Related Modules

- `fxConfig`: Configuration management
- `fxGorm`: Database integration
- Other UTOL modules for complete application setup 