// Dependency registration for the API application.
package internal

import (
	"github.com/UTOL-s/module/api_rest/internal/handler"
	"github.com/UTOL-s/module/api_rest/internal/repository"
	"github.com/UTOL-s/module/api_rest/internal/service"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
	"go.uber.org/fx"
)

func RegisterAll() fx.Option {
	return fx.Options(
		fx.Provide(
			repository.NewUserRepository,
			service.NewUserService,
			handler.NewUserHandler,
			handler.NewHealthHandler,
			handler.NewAuthHandler,
			NewRouter,
			fxSupertoken.AsSuperTokensMiddleware(),
			fxSupertoken.AsVerifySessionMiddleware(),
		),
	)
}
