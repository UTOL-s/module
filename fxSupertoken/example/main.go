package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"

	fxConfig "github.com/UTOL-s/module/fxConfig"
	fxEcho "github.com/UTOL-s/module/fxEcho"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
)

// NewConfig creates a new configuration provider
func NewConfig() *fxConfig.Config {
	config := &fxConfig.Config{}
	config.Accessor = fxConfig.ConfigAccessor()
	return config
}

// NewLogger creates a new logger provider
func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

// Example protected route handler
func ProtectedHandler(c echo.Context) error {
	session := c.Get("supertokensSession")
	return c.JSON(200, map[string]interface{}{
		"message": "This is a protected route",
		"session": session,
	})
}

// Example route that uses session verification
func VerifySessionHandler(c echo.Context) error {
	session := c.Get("session")
	return c.JSON(200, map[string]interface{}{
		"message": "Session verified successfully",
		"session": session,
	})
}

func main() {
	app := fx.New(
		// Provide core dependencies
		fx.Provide(
			NewConfig,
			NewLogger,
		),

		// Register SuperTokens middlewares with fxEcho
		fx.Provide(
			fxSupertoken.AsSuperTokensMiddleware(),
			fxSupertoken.AsVerifySessionMiddleware(),
		),

		// Include the fxSupertoken module
		fxSupertoken.FxSupertoken,

		// Include the fxEcho module
		fxEcho.FxEcho,

		// Register example routes
		fx.Invoke(func(e *echo.Echo) {
			// Protected route using SuperTokens middleware
			e.GET("/protected", ProtectedHandler, fxSupertoken.SupertokenMiddleware)

			// Route using session verification
			e.GET("/verify-session", VerifySessionHandler, fxSupertoken.VerifySession)
		}),

		// Start the application
		fx.Invoke(func(e *echo.Echo, logger *zap.Logger) {
			logger.Info("fxSupertoken with fxEcho example application started",
				zap.String("address", e.Server.Addr),
			)
		}),
	)

	app.Run()
}
