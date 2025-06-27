package fxconfig

import "go.uber.org/fx"

var FxConfig = fx.Module(
	"fxconfig",
	fx.Provide(
		NewConfig,
		NewConfigAccessor,
	),
)
