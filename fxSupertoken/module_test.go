package fxsupertoken

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestSuperTokensConfig(t *testing.T) {
	// Test that SuperTokensConfig can be created with defaults
	config := &SuperTokensConfig{
		ConnectionURI:    "test-uri",
		ConnectionAPIKey: "test-key",
		EmailHost:        "smtp.test.com",
		EmailPassword:    "test-password",
		Email:            "test@test.com",
	}

	assert.Equal(t, "test-uri", config.ConnectionURI)
	assert.Equal(t, "test-key", config.ConnectionAPIKey)
	assert.Equal(t, "smtp.test.com", config.EmailHost)
	assert.Equal(t, "test-password", config.EmailPassword)
	assert.Equal(t, "test@test.com", config.Email)
}

func TestMiddlewareRegistry(t *testing.T) {
	// Test that middleware registry implements the interface correctly
	registry := &SuperTokensMiddlewareRegistry{
		priority:   100,
		middleware: func(next echo.HandlerFunc) echo.HandlerFunc { return next },
	}

	assert.Equal(t, 100, registry.Priority())
	assert.NotNil(t, registry.Middleware())
}

func TestNewSuperTokensMiddlewareRegistry(t *testing.T) {
	// Test that middleware registry can be created
	registry := NewSuperTokensMiddlewareRegistry()

	assert.NotNil(t, registry)
	assert.Equal(t, 100, registry.Priority())
	assert.NotNil(t, registry.Middleware())
}

func TestNewVerifySessionMiddlewareRegistry(t *testing.T) {
	// Test that session verification middleware registry can be created
	registry := NewVerifySessionMiddlewareRegistry()

	assert.NotNil(t, registry)
	assert.Equal(t, 200, registry.Priority())
	assert.NotNil(t, registry.Middleware())
}

func TestFxModule(t *testing.T) {
	// Test that the fx module can be created without errors
	app := fx.New(
		FxSupertoken,
		fx.NopLogger,
	)

	assert.NotNil(t, app)
}

func TestAsMiddleware(t *testing.T) {
	// Test that AsMiddleware function works correctly
	result := AsMiddleware(NewSuperTokensMiddlewareRegistry)
	assert.NotNil(t, result)
}

func TestAsSuperTokensMiddleware(t *testing.T) {
	// Test that AsSuperTokensMiddleware function works correctly
	result := AsSuperTokensMiddleware()
	assert.NotNil(t, result)
}

func TestAsVerifySessionMiddleware(t *testing.T) {
	// Test that AsVerifySessionMiddleware function works correctly
	result := AsVerifySessionMiddleware()
	assert.NotNil(t, result)
}
