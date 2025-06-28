// Route registration for the API application.
package internal

import (
	"github.com/UTOL-s/module/api_rest/internal/handler"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Router handles route registration
type Router struct {
	echo          *echo.Echo
	logger        *zap.Logger
	userHandler   *handler.UserHandler
	authHandler   *handler.AuthHandler
	healthHandler *handler.HealthHandler
}

// NewRouter creates a new router
func NewRouter(
	echo *echo.Echo,
	logger *zap.Logger,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
) *Router {
	return &Router{
		echo:          echo,
		logger:        logger,
		userHandler:   userHandler,
		authHandler:   authHandler,
		healthHandler: healthHandler,
	}
}

// Register registers all routes
func (r *Router) Register() {
	// Welcome route
	r.echo.GET("/", r.welcomeHandler)

	// API routes
	api := r.echo.Group("/api")

	// Health routes
	health := api.Group("/health")
	health.GET("", r.healthHandler.HealthCheck)
	health.GET("/ready", r.healthHandler.ReadinessCheck)

	// Auth routes
	auth := api.Group("/auth")
	auth.GET("/status", r.authHandler.AuthStatus)
	auth.GET("/protected", r.authHandler.ProtectedRoute, fxSupertoken.SupertokenMiddleware)
	auth.GET("/verify", r.authHandler.VerifySession, fxSupertoken.VerifySession)

	// User routes
	users := api.Group("/users")
	users.POST("", r.userHandler.CreateUser)
	users.GET("", r.userHandler.ListUsers)
	users.GET("/search", r.userHandler.SearchUsers)
	users.GET("/:id", r.userHandler.GetUser)
	users.PUT("/:id", r.userHandler.UpdateUser, fxSupertoken.SupertokenMiddleware)
	users.DELETE("/:id", r.userHandler.DeleteUser, fxSupertoken.SupertokenMiddleware)
}

// welcomeHandler handles the welcome route
func (r *Router) welcomeHandler(c echo.Context) error {
	r.logger.Info("welcome route accessed")
	return c.JSON(200, map[string]interface{}{
		"message": "Welcome to Unified Transport Operations League API",
		"version": "1.0.0",
		"status":  "running",
		"endpoints": map[string]interface{}{
			"health": "/api/health",
			"auth":   "/api/auth",
			"users":  "/api/users",
		},
	})
}
