# fxSupertoken

A SuperTokens integration module for fxEcho that provides authentication and session management capabilities.

## Features

- **fxEcho Compatibility**: Fully compatible with the fxEcho module architecture
- **Dependency Injection**: Uses Uber's fx for dependency injection
- **Configuration Management**: Integrates with fxConfig for centralized configuration
- **Middleware Support**: Provides both SuperTokens middleware and session verification middleware
- **Priority-based Middleware**: Supports priority-based middleware registration

## Installation

```bash
go get github.com/UTOL-s/module/fxSupertoken
```

## Configuration

Add SuperTokens configuration to your `config.yaml`:

```yaml
supertokens:
  connection_uri: "your_supertokens_connection_uri"
  connection_api_key: "your_supertokens_api_key"
  app_name: "YourApp"
  api_domain: "http://localhost:8080"
  website_domain: "http://localhost:3000"
  api_base_path: "/api/auth"
  web_base_path: "/api/auth"
  email:
    host: "smtp.gmail.com"
    password: "your_email_password"
    email: "your_email@gmail.com"
```

## Usage

### Basic Integration

```go
package main

import (
    "github.com/labstack/echo/v4"
    "go.uber.org/fx"
    "go.uber.org/zap"

    fxConfig "github.com/UTOL-s/module/fxConfig"
    fxEcho "github.com/UTOL-s/module/fxEcho"
    fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
)

func main() {
    app := fx.New(
        // Provide core dependencies
        fx.Provide(
            NewConfig,
            NewLogger,
        ),

        // Register SuperTokens middlewares
        fx.Provide(
            fxSupertoken.AsSuperTokensMiddleware(),
            fxSupertoken.AsVerifySessionMiddleware(),
        ),

        // Include modules
        fxSupertoken.FxSupertoken,
        fxEcho.FxEcho,

        // Start application
        fx.Invoke(func(e *echo.Echo, logger *zap.Logger) {
            logger.Info("Application started")
        }),
    )

    app.Run()
}
```

### Using Middlewares

#### SuperTokens Middleware

The SuperTokens middleware provides basic session management:

```go
func ProtectedHandler(c echo.Context) error {
    session := c.Get("supertokensSession")
    return c.JSON(200, map[string]interface{}{
        "message": "Protected route",
        "session": session,
    })
}

// Register route with middleware
e.GET("/protected", ProtectedHandler, fxSupertoken.SupertokenMiddleware)
```

#### Session Verification Middleware

The session verification middleware provides enhanced session verification:

```go
func VerifySessionHandler(c echo.Context) error {
    session := c.Get("session")
    return c.JSON(200, map[string]interface{}{
        "message": "Session verified",
        "session": session,
    })
}

// Register route with middleware
e.GET("/verify-session", VerifySessionHandler, fxSupertoken.VerifySession)
```

### Advanced Usage with Route Groups

```go
func main() {
    app := fx.New(
        // ... other providers

        fx.Invoke(func(e *echo.Echo) {
            // Create protected route group
            protected := e.Group("/api")
            protected.Use(fxSupertoken.SupertokenMiddleware)
            
            protected.GET("/users", GetUsersHandler)
            protected.POST("/users", CreateUserHandler)
            
            // Create admin route group with session verification
            admin := e.Group("/admin")
            admin.Use(fxSupertoken.VerifySession)
            
            admin.GET("/dashboard", AdminDashboardHandler)
        }),
    )

    app.Run()
}
```

## API Reference

### Types

#### SuperTokensConfig

Configuration structure for SuperTokens:

```go
type SuperTokensConfig struct {
    ConnectionURI    string
    ConnectionAPIKey string
    EmailHost        string
    EmailPassword    string
    Email            string
    APIBasePath      string
    WebBasePath      string
    AppName          string
    APIDomain        string
    WebsiteDomain    string
}
```

#### MiddlewareRegistryIf

Interface for middleware registration:

```go
type MiddlewareRegistryIf interface {
    Priority() int
    Middleware() echo.MiddlewareFunc
}
```

### Functions

#### NewSuperTokensConfig

Creates SuperTokens configuration from fxConfig:

```go
func NewSuperTokensConfig(config *fxConfig.Config) (*SuperTokensConfig, error)
```

#### NewSuperTokensMiddleware

Creates the main SuperTokens middleware:

```go
func NewSuperTokensMiddleware() echo.MiddlewareFunc
```

#### NewVerifySessionMiddleware

Creates the session verification middleware:

```go
func NewVerifySessionMiddleware() echo.MiddlewareFunc
```

#### AsSuperTokensMiddleware

Annotates SuperTokens middleware for fxEcho:

```go
func AsSuperTokensMiddleware() any
```

#### AsVerifySessionMiddleware

Annotates session verification middleware for fxEcho:

```go
func AsVerifySessionMiddleware() any
```

### Module

#### FxSupertoken

The main fx module that provides all SuperTokens functionality:

```go
var FxSupertoken = fx.Module(
    "fxsupertoken",
    fx.Provide(
        NewSuperTokensConfig,
        NewSuperTokensMiddleware,
        NewVerifySessionMiddleware,
        NewSuperTokensMiddlewareRegistry,
        NewVerifySessionMiddlewareRegistry,
    ),
    fx.Invoke(InitSuperTokens),
)
```

## Middleware Priority

The module uses priority-based middleware registration:

- **SuperTokens Middleware**: Priority 100
- **Session Verification Middleware**: Priority 200

Higher priority middlewares are executed first.

## Error Handling

The module provides proper error handling for:

- Configuration loading errors
- SuperTokens initialization errors
- Session verification errors
- Middleware execution errors

## Examples

See the `example/` directory for complete working examples.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License. 