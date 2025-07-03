package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"

	fxEcho "github.com/UTOL-s/module/fxEcho"
	"github.com/UTOL-s/module/fxEcho/example/handlers"
	"github.com/UTOL-s/module/fxEcho/example/middleware"
	"github.com/UTOL-s/module/fxEcho/example/models"
	"github.com/UTOL-s/module/fxEcho/example/providers"
	"github.com/UTOL-s/module/fxEcho/example/routes"
)

func main() {
	app := fx.New(
		// Provide core dependencies
		fx.Provide(
			providers.NewConfig,
			providers.NewLogger,
			models.NewUserService,
			handlers.NewUserHandler,
			handlers.NewHealthHandler,
			handlers.NewWelcomeHandler,
			middleware.NewRequestTimingMiddleware,
			middleware.NewRequestIDMiddleware,
		),

		// Register routes using the builder pattern
		fx.Provide(
			fxEcho.AsRoute(routes.NewWelcomeRoute),
			fxEcho.AsRoute(routes.NewHealthRoute),
			fxEcho.AsGroup(routes.NewUserRoutes),
			fxEcho.AsGroup(routes.NewAdminRoutes),
		),

		// Include the fxEcho module
		fxEcho.FxEcho,

		// Start the application
		fx.Invoke(func(e *echo.Echo, logger *zap.Logger) {
			logger.Info("fxEcho example application started",
				zap.String("address", e.Server.Addr),
			)
		}),
	)

	app.Run()
}
