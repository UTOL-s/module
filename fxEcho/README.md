# FX Echo Module

A highly optimized module for integrating the Echo web framework with Uber FX dependency injection in the UTOL module system.

## Overview

The fxEcho module provides seamless integration between the Echo web framework and the FX dependency injection system. This module handles Echo server setup, middleware configuration, route management, and lifecycle management through dependency injection with enhanced performance and reliability features.

## Features

✅ **Echo Server Setup**: Automatic Echo server initialization with FX
✅ **Middleware Integration**: Dependency injection for Echo middleware with priority support
✅ **Route Management**: FX-based route registration with fluent builder API
✅ **Configuration Integration**: Integration with fxConfig for server settings
✅ **Graceful Shutdown**: Proper server shutdown handling with timeout management
✅ **Health Checks**: Built-in health check endpoints
✅ **CORS Support**: Configurable CORS middleware
✅ **Request Logging**: Structured request logging middleware
✅ **Error Handling**: Comprehensive error handling and logging
✅ **Performance Optimizations**: HTTP timeouts, connection pooling, and efficient routing
✅ **Type Safety**: Strongly typed route and group builders
✅ **Testing Support**: Comprehensive test utilities and examples

## Quick Start

```go
package main

import (
    "net/http"
    
    "github.com/UTOL-s/module/fxEcho"
    "github.com/UTOL-s/module/fxConfig"
    "github.com/labstack/echo/v4"
    "go.uber.org/fx"
    "go.uber.org/zap"
)

func main() {
    app := fx.New(
        fxConfig.FxConfig,
        fx.Provide(
            func() *zap.Logger {
                return zap.NewNop()
            },
            // Define routes using the fluent builder API
            func() fxEcho.RouteRegistryIf {
                return fxEcho.GET("/hello", func(c echo.Context) error {
                    return c.JSON(http.StatusOK, map[string]string{
                        "message": "Hello World",
                    })
                }).Build()
            },
            // Define route groups
            func() fxEcho.GroupRegistryIf {
                return fxEcho.NewGroup("/api/v1").
                    AddRoute(fxEcho.GET("/users", func(c echo.Context) error {
                        return c.JSON(http.StatusOK, map[string]interface{}{
                            "users": []string{"user1", "user2"},
                        })
                    }).Build()).
                    AddRoute(fxEcho.POST("/users", func(c echo.Context) error {
                        return c.JSON(http.StatusCreated, map[string]string{
                            "message": "User created",
                        })
                    }).Build()).
                    Build()
            },
        ),
        fx.Annotate(
            func() fxEcho.RouteRegistryIf {
                return fxEcho.GET("/hello", func(c echo.Context) error {
                    return c.JSON(http.StatusOK, map[string]string{
                        "message": "Hello World",
                    })
                }).Build()
            },
            fxEcho.AsRoute,
        ),
        fx.Annotate(
            func() fxEcho.GroupRegistryIf {
                return fxEcho.NewGroup("/api/v1").
                    AddRoute(fxEcho.GET("/users", func(c echo.Context) error {
                        return c.JSON(http.StatusOK, map[string]interface{}{
                            "users": []string{"user1", "user2"},
                        })
                    }).Build()).
                    Build()
            },
            fxEcho.AsGroup,
        ),
        fxEcho.FxEcho,
        fx.Invoke(func(e *echo.Echo) {
            // Echo server will be automatically started
        }),
    )
    
    app.Run()
}
```

## Configuration

The module supports comprehensive configuration through the fxConfig module:

```yaml
server:
  host: "0.0.0.0"
  port: "8080"
  read_timeout: 30
  write_timeout: 30
  idle_timeout: 60

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

### Default Values

- **Host**: `0.0.0.0`
- **Port**: `8080`
- **Read Timeout**: `30 seconds`
- **Write Timeout**: `30 seconds`
- **Idle Timeout**: `60 seconds`

## API Reference

### Route Builder API

```go
// Create routes with fluent API
route := fxEcho.GET("/users", handler).Build()
route := fxEcho.POST("/users", handler).Build()
route := fxEcho.PUT("/users/:id", handler).Build()
route := fxEcho.DELETE("/users/:id", handler).Build()
route := fxEcho.PATCH("/users/:id", handler).Build()

// Custom route
route := fxEcho.NewRoute("OPTIONS", "/users", handler).Build()
```

### Group Builder API

```go
// Create route groups
group := fxEcho.NewGroup("/api/v1").
    AddRoute(fxEcho.GET("/users", handler).Build()).
    AddRoute(fxEcho.POST("/users", handler).Build()).
    Build()

// Nested groups
parentGroup := fxEcho.NewGroup("/api").
    AddGroup(fxEcho.NewGroup("/v1").
        AddRoute(fxEcho.GET("/health", handler).Build()).
        Build()).
    Build()
```

### Middleware Integration

```go
// Provide custom middleware
func NewCustomMiddleware() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Response().Header().Set("X-Custom-Header", "value")
            return next(c)
        }
    }
}

// Use with FX annotation
fx.Annotate(
    NewCustomMiddleware,
    fx.ResultTags(`group:"middlewares"`),
)
```

## Built-in Features

### Health Check Endpoint

The module automatically provides a health check endpoint at `/health`:

```json
{
  "status": "healthy",
  "time": "2024-01-01T00:00:00Z"
}
```

### Default Middleware

When no custom middlewares are provided, the module automatically includes:

- **Logger**: Request logging
- **Recover**: Panic recovery
- **CORS**: Cross-origin resource sharing

### Graceful Shutdown

The module implements graceful shutdown with:

- 30-second shutdown timeout
- Proper connection cleanup
- Structured logging during shutdown

## Performance Optimizations

1. **HTTP Timeouts**: Configurable read, write, and idle timeouts
2. **Connection Pooling**: Efficient connection management
3. **Middleware Optimization**: Priority-based middleware ordering
4. **Route Caching**: Optimized route registration
5. **Memory Management**: Proper resource cleanup

## Error Handling

The module provides comprehensive error handling:

- Configuration validation
- Server startup error handling
- Graceful shutdown error handling
- Structured logging with zap
- Health check endpoint for monitoring

## Testing

The module includes comprehensive test utilities:

```go
func TestFxEchoModule(t *testing.T) {
    app := fxtest.New(t,
        fx.Provide(
            newTestConfig,
            newTestLogger,
            NewExampleRoute,
        ),
        fx.Annotate(NewExampleRoute, fxEcho.AsRoute),
        fxEcho.FxEcho,
    )

    app.RequireStart()
    defer app.RequireStop()

    var e *echo.Echo
    err := app.Invoke(func(echo *echo.Echo) {
        e = echo
    })
    assert.NoError(t, err)
    assert.NotNil(t, e)
}
```

## Dependencies

- `github.com/labstack/echo/v4`: Echo web framework
- `go.uber.org/fx`: Dependency injection framework
- `go.uber.org/zap`: Structured logging
- `github.com/stretchr/testify`: Testing utilities

## Migration from Previous Version

The optimized version includes breaking changes for better performance and maintainability:

1. **Error Returns**: `NewEcho` now returns an error
2. **Logger Dependency**: Logger is now required in `EchoParams`
3. **Server Configuration**: New `ServerConfig` struct for better configuration management
4. **Builder API**: New fluent builder API for routes and groups
5. **Enhanced Middleware**: Better middleware integration with priority support

## Contributing

This module is open for contributions. Please check the main project repository for contribution guidelines.

## Related Modules

- `fxConfig`: Configuration management
- `fxGorm`: Database integration
- Other UTOL modules for complete application setup 

## Example Usage: Handler and Handler Group

Below is a minimal example showing how to register a single handler and a handler group using the builder API and provide them to an `fx.App`:

```go
package main

import (
    "net/http"
    "github.com/UTOL-s/module/fxEcho"
    "github.com/labstack/echo/v4"
    "go.uber.org/fx"
    "go.uber.org/zap"
)

// Single handler
func helloHandler(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{"message": "Hello from handler!"})
}

// Group handler
func userListHandler(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{"users": []string{"alice", "bob"}})
}

func main() {
    app := fx.New(
        fx.Provide(
            // Logger
            func() *zap.Logger { return zap.NewNop() },
            // Register a single route handler
            fxEcho.AsRoute(func() fxEcho.RouteRegistryIf {
                return fxEcho.GET("/hello", helloHandler).Build()
            }),
            // Register a group of handlers
            fxEcho.AsGroup(func() fxEcho.GroupRegistryIf {
                return fxEcho.NewGroup("/api").
                    AddRoute(fxEcho.GET("/users", userListHandler).Build()).
                    Build()
            }),
        ),
        fxEcho.FxEcho,
        fx.Invoke(func(e *echo.Echo) {
            // Echo server will be started automatically
        }),
    )
    app.Run()
}
```

This will expose:
- `GET /hello` → returns `{ "message": "Hello from handler!" }`
- `GET /api/users` → returns `{ "users": ["alice", "bob"] }` 