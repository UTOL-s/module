package FxEcho

import (
	"context"
	"fmt"
	"net/http"
	"time"

	fxConfig "github.com/UTOL-s/module/fxConfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "fxecho"

// ServerConfig holds Echo server configuration
type ServerConfig struct {
	Host         string        `mapstructure:"host"`
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// EchoParams holds all dependencies for Echo server
type EchoParams struct {
	fx.In
	Lifecycle   fx.Lifecycle
	Routes      []RouteRegistryIf `group:"routes"`
	Groups      []GroupRegistryIf `group:"groups"`
	Config      *fxConfig.Config
	Middlewares []echo.MiddlewareFunc `group:"middlewares"`
	Logger      *zap.Logger
}

var FxEcho = fx.Module(
	ModuleName,
	fx.Provide(
		NewEcho,
		NewServerConfig,
	),
	fx.Invoke(func(e *echo.Echo) {}),
)

// NewServerConfig creates server configuration from fxConfig
func NewServerConfig(config *fxConfig.Config) (*ServerConfig, error) {
	serverConfig := &ServerConfig{
		Host:         config.Accessor.String("server.host"),
		Port:         config.Accessor.String("server.port"),
		ReadTimeout:  time.Duration(config.Accessor.Int("server.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(config.Accessor.Int("server.write_timeout")) * time.Second,
		IdleTimeout:  time.Duration(config.Accessor.Int("server.idle_timeout")) * time.Second,
	}

	// Set defaults if not configured
	if serverConfig.Host == "" {
		serverConfig.Host = "0.0.0.0"
	}
	if serverConfig.Port == "" {
		serverConfig.Port = "8080"
	}
	if serverConfig.ReadTimeout == 0 {
		serverConfig.ReadTimeout = 30 * time.Second
	}
	if serverConfig.WriteTimeout == 0 {
		serverConfig.WriteTimeout = 30 * time.Second
	}
	if serverConfig.IdleTimeout == 0 {
		serverConfig.IdleTimeout = 60 * time.Second
	}

	return serverConfig, nil
}

// NewEcho creates and configures Echo server with FX lifecycle management
func NewEcho(p EchoParams) (*echo.Echo, error) {
	// Create Echo instance
	e := echo.New()
	e.HideBanner = true

	// Configure server settings
	serverConfig, err := NewServerConfig(p.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to create server config: %w", err)
	}

	// Set up HTTP server with timeouts
	e.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port),
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
		IdleTimeout:  serverConfig.IdleTimeout,
	}

	// Add default middlewares if none provided
	if len(p.Middlewares) == 0 {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}))
	} else {
		// Add custom middlewares
		for _, m := range p.Middlewares {
			e.Use(m)
		}
	}

	// Register route groups
	for _, group := range p.Groups {
		g := e.Group(group.Prefix())
		group.Register(g)
	}

	// Register individual routes
	for _, route := range p.Routes {
		e.Add(route.Method(), route.Path(), route.Handle)
	}

	// Add health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "healthy",
			"time":   time.Now().UTC(),
		})
	})

	// Configure FX lifecycle hooks
	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			p.Logger.Info("starting Echo server",
				zap.String("address", e.Server.Addr),
				zap.String("host", serverConfig.Host),
				zap.String("port", serverConfig.Port),
			)

			go func() {
				if err := e.Start(e.Server.Addr); err != nil && err != http.ErrServerClosed {
					p.Logger.Error("failed to start Echo server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			p.Logger.Info("shutting down Echo server")

			// Create shutdown context with timeout
			shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := e.Shutdown(shutdownCtx); err != nil {
				p.Logger.Error("error during server shutdown", zap.Error(err))
				return err
			}

			p.Logger.Info("Echo server shutdown completed")
			return nil
		},
	})

	return e, nil
}
