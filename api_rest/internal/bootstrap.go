// Application bootstrap and lifecycle management.
package internal

import (
	fxConfig "github.com/UTOL-s/module/fxConfig"
	fxEcho "github.com/UTOL-s/module/fxEcho"
	fxGorm "github.com/UTOL-s/module/fxGorm"
	fxSupertoken "github.com/UTOL-s/module/fxSupertoken"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Bootstrap() *fx.App {
	return fx.New(
		fx.Provide(func() (*zap.Logger, error) {
			return zap.NewProduction()
		}),
		fxConfig.FxConfig,
		fxEcho.FxEcho,
		fxGorm.FxGorm,
		fxSupertoken.FxSupertoken,
		RegisterAll(),
		fx.Provide(func(config *fxConfig.Config) fxGorm.Params {
			return fxGorm.Params{Config: config}
		}),
		fx.Invoke(func(router *Router) { router.Register() }),
	)
}
