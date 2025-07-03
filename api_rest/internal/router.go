// Route registration for the API application.
package internal

import (
	"github.com/UTOL-s/module/api_rest/internal/handler"
	fxConfig "github.com/UTOL-s/module/fxConfig"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Router handles route registration
type Router struct {
	echo              *echo.Echo
	logger            *zap.Logger
	config            *fxConfig.Config
	userHandler       *handler.UserHandler
	healthHandler     *handler.HealthHandler
	superTokensConfig *fxSupertoken.SuperTokensConfig
}

// NewRouter creates a new router
func NewRouter(
	echo *echo.Echo,
	logger *zap.Logger,
	config *fxConfig.Config,
	userHandler *handler.UserHandler,
	healthHandler *handler.HealthHandler,
	superTokensConfig *fxSupertoken.SuperTokensConfig,
) *Router {
	return &Router{
		echo:              echo,
		logger:            logger,
		config:            config,
		userHandler:       userHandler,
		healthHandler:     healthHandler,
		superTokensConfig: superTokensConfig,
	}
}

// Register registers all routes
func (r *Router) Register() {
	// Welcome route
	r.echo.GET("/", r.welcomeHandler)

	// Apply SuperTokens middleware globally to handle auth endpoints
	if r.superTokensConfig.IsInitialized {
		r.echo.Use(fxSupertoken.SuperTokensMiddlewareWrapper())
	}

	// API routes
	api := r.echo.Group("/api")

	// Health routes
	health := api.Group("/health")
	health.GET("", r.healthHandler.HealthCheck)
	health.GET("/ready", r.healthHandler.ReadinessCheck)

	// User routes
	users := api.Group("/v1/users")
	users.POST("", r.userHandler.CreateUser)
	users.GET("", r.userHandler.ListUsers)
	users.GET("/search", r.userHandler.SearchUsers)
	users.GET("/:id", r.userHandler.GetUser, fxSupertoken.VerifySession(r.superTokensConfig))
	users.PUT("/:id", r.userHandler.UpdateUser, fxSupertoken.VerifySession(r.superTokensConfig))
	users.DELETE("/:id", r.userHandler.DeleteUser, fxSupertoken.SupertokenMiddleware(r.superTokensConfig))
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
