package routes

import (
	fxEcho "github.com/UTOL-s/module/fxEcho"
	"github.com/UTOL-s/module/fxEcho/example/handlers"
	"github.com/labstack/echo/v4"
)

// NewWelcomeRoute creates the welcome route
func NewWelcomeRoute(welcomeHandler *handlers.WelcomeHandler) fxEcho.RouteRegistryIf {
	return fxEcho.GET("/", welcomeHandler.Welcome).Build()
}

// NewHealthRoute creates the health route
func NewHealthRoute(healthHandler *handlers.HealthHandler) fxEcho.RouteRegistryIf {
	return fxEcho.GET("/health", healthHandler.Health).Build()
}

// NewUserRoutes creates the user routes group with middleware
func NewUserRoutes(userHandler *handlers.UserHandler, requestTimingMiddleware echo.MiddlewareFunc, requestIDMiddleware echo.MiddlewareFunc) fxEcho.GroupRegistryIf {
	return fxEcho.NewGroup("/api").
		Use(requestTimingMiddleware, requestIDMiddleware). // Apply middleware to the entire API group
		AddRoute(fxEcho.GET("/users", userHandler.ListUsers).Build()).
		AddRoute(fxEcho.GET("/users/:id", userHandler.GetUser).Build()).
		AddRoute(fxEcho.POST("/users", userHandler.CreateUser).Build()).
		Build()
}

// NewAdminRoutes creates admin routes with different middleware (example for future use)
func NewAdminRoutes(requestTimingMiddleware echo.MiddlewareFunc) fxEcho.GroupRegistryIf {
	return fxEcho.NewGroup("/admin").
		Use(requestTimingMiddleware). // Only timing middleware for admin routes
		AddRoute(fxEcho.GET("/dashboard", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"message": "Admin dashboard"})
		}).Build()).
		Build()
}
