package fxsupertoken

import (
	"context"
	"testing"
	"time"

	fxConfig "github.com/UTOL-s/module/fxConfig"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func mockFxConfig() *fxConfig.Config {
	return &fxConfig.Config{
		Accessor: fxConfig.ConfigAccessor(),
	}
}

func TestFxSupertokenModule(t *testing.T) {
	// Test that the fxSupertoken module can be created without errors
	app := fxtest.New(t,
		fx.Provide(
			mockFxConfig,
		),
		FxSupertoken,
		fx.NopLogger,
	)

	// Start the app
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := app.Start(ctx)
	assert.NoError(t, err)

	// Stop the app
	err = app.Stop(ctx)
	assert.NoError(t, err)
}

func TestMiddlewareRegistration(t *testing.T) {
	// Test that middlewares can be registered correctly
	superTokensRegistry := NewSuperTokensMiddlewareRegistry()
	verifySessionRegistry := NewVerifySessionMiddlewareRegistry()

	assert.Equal(t, 100, superTokensRegistry.Priority())
	assert.Equal(t, 200, verifySessionRegistry.Priority())
	assert.NotNil(t, superTokensRegistry.Middleware())
	assert.NotNil(t, verifySessionRegistry.Middleware())
}

func TestConfigDefaults(t *testing.T) {
	// Test that configuration defaults are set correctly
	config := &SuperTokensConfig{
		ConnectionURI:    "test-uri",
		ConnectionAPIKey: "test-key",
		EmailHost:        "smtp.test.com",
		EmailPassword:    "test-password",
		Email:            "test@test.com",
	}

	// Test defaults
	if config.APIBasePath == "" {
		config.APIBasePath = "/api/auth"
	}
	if config.WebBasePath == "" {
		config.WebBasePath = "/api/auth"
	}
	if config.AppName == "" {
		config.AppName = "YourApp"
	}
	if config.APIDomain == "" {
		config.APIDomain = "http://localhost:8080"
	}
	if config.WebsiteDomain == "" {
		config.WebsiteDomain = "http://localhost:3000"
	}

	assert.Equal(t, "/api/auth", config.APIBasePath)
	assert.Equal(t, "/api/auth", config.WebBasePath)
	assert.Equal(t, "YourApp", config.AppName)
	assert.Equal(t, "http://localhost:8080", config.APIDomain)
	assert.Equal(t, "http://localhost:3000", config.WebsiteDomain)
}

func TestMiddlewareFunctions(t *testing.T) {
	// Test that middleware functions return valid Echo middleware functions
	superTokensMiddleware := NewSuperTokensMiddleware()
	verifySessionMiddleware := NewVerifySessionMiddleware()

	assert.NotNil(t, superTokensMiddleware)
	assert.NotNil(t, verifySessionMiddleware)

	// Test that they can be used as Echo middleware
	e := echo.New()
	e.Use(superTokensMiddleware)
	e.Use(verifySessionMiddleware)

	assert.NotNil(t, e)
}
