package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	logger *zap.Logger
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(logger *zap.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

// Health handles GET /health
func (h *HealthHandler) Health(c echo.Context) error {
	h.logger.Debug("health check requested")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "fxEcho Example",
	})
}
