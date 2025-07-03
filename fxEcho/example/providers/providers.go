package providers

import (
	fxConfig "github.com/UTOL-s/module/fxConfig"
	"go.uber.org/zap"
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
