// Dependency registration for the API application.
package internal

import (
	"github.com/UTOL-s/module/api_rest/internal/handler"
	"github.com/UTOL-s/module/api_rest/internal/repository"
	"github.com/UTOL-s/module/api_rest/internal/service"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterAll() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.NewUserRepository,
			service.NewUserService,
			handler.NewUserHandler,
			handler.NewHealthHandler,
			NewRouter,
			// Provide CORS middleware through FX system
			fx.Annotate(
				func() echo.MiddlewareFunc {
					return fxSupertoken.DefaultCORSHandler()
				},
				fx.ResultTags(`group:"middlewares"`),
			),
			fxSupertoken.AsSuperTokensMiddleware(),
			fxSupertoken.AsVerifySessionMiddleware(),
			fxSupertoken.AsSuperTokensWrapperMiddleware(),
		),
	)
}
