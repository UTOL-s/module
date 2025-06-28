package fxsupertoken

import (
	fxConfig "github.com/UTOL-s/module/fxConfig"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const ModuleName = "fxsupertoken"

// MiddlewareRegistryIf defines the interface for middleware registration
type MiddlewareRegistryIf interface {
	Priority() int
	Middleware() echo.MiddlewareFunc
}

// SuperTokensConfig holds SuperTokens configuration
type SuperTokensConfig struct {
	ConnectionURI    string `mapstructure:"connection_uri"`
	ConnectionAPIKey string `mapstructure:"connection_api_key"`
	EmailHost        string `mapstructure:"email.host"`
	EmailPassword    string `mapstructure:"email.password"`
	Email            string `mapstructure:"email"`
	APIBasePath      string `mapstructure:"api_base_path"`
	WebBasePath      string `mapstructure:"web_base_path"`
	AppName          string `mapstructure:"app_name"`
	APIDomain        string `mapstructure:"api_domain"`
	WebsiteDomain    string `mapstructure:"website_domain"`
}

// NewSuperTokensConfig creates SuperTokens configuration from fxConfig
func NewSuperTokensConfig(config *fxConfig.Config) (*SuperTokensConfig, error) {
	superTokensConfig := &SuperTokensConfig{
		ConnectionURI:    config.Accessor.String("supertokens.connection_uri"),
		ConnectionAPIKey: config.Accessor.String("supertokens.connection_api_key"),
		EmailHost:        config.Accessor.String("supertokens.email.host"),
		EmailPassword:    config.Accessor.String("supertokens.email.password"),
		Email:            config.Accessor.String("supertokens.email"),
		APIBasePath:      config.Accessor.String("supertokens.api_base_path"),
		WebBasePath:      config.Accessor.String("supertokens.web_base_path"),
		AppName:          config.Accessor.String("supertokens.app_name"),
		APIDomain:        config.Accessor.String("supertokens.api_domain"),
		WebsiteDomain:    config.Accessor.String("supertokens.website_domain"),
	}

	// Set defaults if not configured
	if superTokensConfig.APIBasePath == "" {
		superTokensConfig.APIBasePath = "/api/auth"
	}
	if superTokensConfig.WebBasePath == "" {
		superTokensConfig.WebBasePath = "/api/auth"
	}
	if superTokensConfig.AppName == "" {
		superTokensConfig.AppName = "YourApp"
	}
	if superTokensConfig.APIDomain == "" {
		superTokensConfig.APIDomain = "http://localhost:8080"
	}
	if superTokensConfig.WebsiteDomain == "" {
		superTokensConfig.WebsiteDomain = "http://localhost:3000"
	}

	return superTokensConfig, nil
}

// NewSuperTokensMiddleware creates the main SuperTokens middleware
func NewSuperTokensMiddleware() echo.MiddlewareFunc {
	return SupertokenMiddleware
}

// NewVerifySessionMiddleware creates the session verification middleware
func NewVerifySessionMiddleware() echo.MiddlewareFunc {
	return VerifySession
}

// SuperTokensMiddlewareRegistry implements MiddlewareRegistryIf for SuperTokens middleware
type SuperTokensMiddlewareRegistry struct {
	priority   int
	middleware echo.MiddlewareFunc
}

func (s *SuperTokensMiddlewareRegistry) Priority() int {
	return s.priority
}

func (s *SuperTokensMiddlewareRegistry) Middleware() echo.MiddlewareFunc {
	return s.middleware
}

// NewSuperTokensMiddlewareRegistry creates a middleware registry for SuperTokens
func NewSuperTokensMiddlewareRegistry() MiddlewareRegistryIf {
	return &SuperTokensMiddlewareRegistry{
		priority:   100, // High priority for auth middleware
		middleware: SupertokenMiddleware,
	}
}

// NewVerifySessionMiddlewareRegistry creates a middleware registry for session verification
func NewVerifySessionMiddlewareRegistry() MiddlewareRegistryIf {
	return &SuperTokensMiddlewareRegistry{
		priority:   200, // Higher priority for session verification
		middleware: VerifySession,
	}
}

// AsMiddleware annotates a middleware constructor for fxEcho compatibility
func AsMiddleware(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(MiddlewareRegistryIf)),
		fx.ResultTags(`group:"middlewares"`),
	)
}

// AsSuperTokensMiddleware annotates SuperTokens middleware for fxEcho
func AsSuperTokensMiddleware() any {
	return AsMiddleware(NewSuperTokensMiddlewareRegistry)
}

// AsVerifySessionMiddleware annotates session verification middleware for fxEcho
func AsVerifySessionMiddleware() any {
	return AsMiddleware(NewVerifySessionMiddlewareRegistry)
}

var FxSupertoken = fx.Module(
	ModuleName,
	fx.Provide(
		NewSuperTokensConfig,
		fx.Annotate(
			NewSuperTokensMiddlewareRegistry,
			fx.ResultTags(`group:"middlewares"`),
		),
		fx.Annotate(
			NewVerifySessionMiddlewareRegistry,
			fx.ResultTags(`group:"middlewares"`),
		),
	),
	fx.Invoke(InitSuperTokens),
)
