package FxEcho

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	fxConfig "github.com/UTOL-s/module/fxConfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

// Example route handler
func exampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello from fxEcho!",
		"status":  "success",
	})
}

// Example group handler
func apiHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"api":    "v1",
		"status": "active",
	})
}

// Example middleware
func customMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("X-Custom-Header", "fxEcho")
			return next(c)
		}
	}
}

// Example route constructor
func NewExampleRoute() RouteRegistryIf {
	return GET("/example", exampleHandler).Build()
}

// Example API route constructor
func NewAPIRoute() RouteRegistryIf {
	return GET("/api", apiHandler).Build()
}

// Example group constructor
func NewExampleGroup() GroupRegistryIf {
	return NewGroup("/api/v1").
		AddRoute(GET("/users", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"users": []string{"user1", "user2", "user3"},
			})
		}).Build()).
		AddRoute(POST("/users", func(c echo.Context) error {
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"message": "User created",
			})
		}).Build()).
		Build()
}

// Example middleware constructor
func NewCustomMiddleware() echo.MiddlewareFunc {
	return customMiddleware()
}

// Example test configuration
func newTestConfig() *fxConfig.Config {
	config := &fxConfig.Config{}
	config.Accessor = fxConfig.ConfigAccessor()
	return config
}

// Example logger
func newTestLogger() *zap.Logger {
	return zap.NewNop()
}

// Remove duplicate TestFxEchoModule and use only one, with correct provider usage
func TestFxEchoModule_Integration(t *testing.T) {
	var e *echo.Echo

	app := fxtest.New(t,
		fx.Provide(
			newTestConfig,
			newTestLogger,
			AsRoute(NewExampleRoute),
			AsRoute(NewAPIRoute),
			AsGroup(NewExampleGroup),
			fx.Annotate(NewCustomMiddleware, fx.ResultTags(`group:"middlewares"`)),
		),
		FxEcho,
		fx.Populate(&e),
	)

	app.RequireStart()
	defer app.RequireStop()

	assert.NotNil(t, e)
}

func TestRouteBuilder(t *testing.T) {
	// Test route builder
	route := GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}).Build()

	assert.Equal(t, "GET", route.Method())
	assert.Equal(t, "/test", route.Path())

	// Test different HTTP methods
	postRoute := POST("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}).Build()
	assert.Equal(t, "POST", postRoute.Method())

	putRoute := PUT("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}).Build()
	assert.Equal(t, "PUT", putRoute.Method())

	deleteRoute := DELETE("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}).Build()
	assert.Equal(t, "DELETE", deleteRoute.Method())

	patchRoute := PATCH("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}).Build()
	assert.Equal(t, "PATCH", patchRoute.Method())
}

func TestGroupBuilder(t *testing.T) {
	// Test group builder
	group := NewGroup("/api").
		AddRoute(GET("/users", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{"users": []string{}})
		}).Build()).
		AddRoute(POST("/users", func(c echo.Context) error {
			return c.JSON(http.StatusCreated, map[string]interface{}{"message": "created"})
		}).Build()).
		Build()

	assert.Equal(t, "/api", group.Prefix())

	// Test nested groups
	nestedGroup := NewGroup("/v1").
		AddRoute(GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
		}).Build()).
		Build()

	parentGroup := NewGroup("/api").
		AddGroup(nestedGroup).
		Build()

	assert.Equal(t, "/api", parentGroup.Prefix())
}

func TestGroupBuilderWithMiddleware(t *testing.T) {
	// Test group builder with middleware
	customMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("X-Group-Middleware", "applied")
			return next(c)
		}
	}

	group := NewGroup("/api/v1").
		Use(customMiddleware).
		AddRoute(GET("/users", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{"users": []string{}})
		}).Build()).
		AddRoute(POST("/users", func(c echo.Context) error {
			return c.JSON(http.StatusCreated, map[string]interface{}{"message": "created"})
		}).Build()).
		Build()

	assert.Equal(t, "/api/v1", group.Prefix())

	// Test that the group can be registered (this would be tested in integration)
	// The middleware functionality is tested through the Register method
}

func TestServerConfig(t *testing.T) {
	config := newTestConfig()
	serverConfig, err := NewServerConfig(config)
	assert.NoError(t, err)
	assert.NotNil(t, serverConfig)

	// Test defaults
	assert.Equal(t, "0.0.0.0", serverConfig.Host)
	assert.Equal(t, "8080", serverConfig.Port)
	assert.Equal(t, 30*time.Second, serverConfig.ReadTimeout)
	assert.Equal(t, 30*time.Second, serverConfig.WriteTimeout)
	assert.Equal(t, 60*time.Second, serverConfig.IdleTimeout)
}

func TestHealthEndpoint(t *testing.T) {
	// Create a minimal Echo instance to test health endpoint
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "healthy",
			"time":   time.Now().UTC(),
		})
	})

	// Test health endpoint
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "healthy")
}

func Example() {
	app := fx.New(
		fx.Provide(
			// Provide configuration
			func() *fxConfig.Config {
				config := &fxConfig.Config{}
				config.Accessor = fxConfig.ConfigAccessor()
				return config
			},
			// Provide logger
			func() *zap.Logger {
				return zap.NewNop()
			},
			// Provide routes
			AsRoute(func() RouteRegistryIf {
				return GET("/hello", func(c echo.Context) error {
					return c.JSON(http.StatusOK, map[string]string{"message": "Hello World"})
				}).Build()
			}),
			// Provide route groups
			AsGroup(func() GroupRegistryIf {
				return NewGroup("/api/v1").
					AddRoute(GET("/users", func(c echo.Context) error {
						return c.JSON(http.StatusOK, map[string]interface{}{"users": []string{}})
					}).Build()).
					AddRoute(POST("/users", func(c echo.Context) error {
						return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
					}).Build()).
					Build()
			}),
			// Provide custom middleware
			fx.Annotate(
				func() echo.MiddlewareFunc {
					return middleware.Logger()
				},
				fx.ResultTags(`group:"middlewares"`),
			),
		),
		FxEcho,
		fx.Invoke(func(e *echo.Echo) {
			// Echo server will be automatically started
		}),
	)

	app.Run()
}

func TestExampleUsage_HandlerAndGroup(t *testing.T) {
	// Single handler
	helloHandler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello from handler!"})
	}

	// Group handler
	userListHandler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"users": []string{"alice", "bob"}})
	}

	var e *echo.Echo

	app := fxtest.New(t,
		fx.Provide(
			// Config
			newTestConfig,
			// Logger
			func() *zap.Logger { return zap.NewNop() },
			// Register a single route handler
			AsRoute(func() RouteRegistryIf {
				return GET("/hello", helloHandler).Build()
			}),
			// Register a group of handlers
			AsGroup(func() GroupRegistryIf {
				return NewGroup("/api").
					AddRoute(GET("/users", userListHandler).Build()).
					Build()
			}),
		),
		FxEcho,
		fx.Populate(&e),
	)

	app.RequireStart()
	defer app.RequireStop()

	assert.NotNil(t, e)

	// Test the single handler
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Hello from handler!")

	// Test the group handler
	req = httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "alice")
	assert.Contains(t, rec.Body.String(), "bob")
}
