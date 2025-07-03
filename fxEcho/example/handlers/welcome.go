package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// WelcomeHandler handles welcome page requests
type WelcomeHandler struct {
	logger *zap.Logger
}

// NewWelcomeHandler creates a new welcome handler
func NewWelcomeHandler(logger *zap.Logger) *WelcomeHandler {
	return &WelcomeHandler{logger: logger}
}

// Welcome handles GET /
func (h *WelcomeHandler) Welcome(c echo.Context) error {
	h.logger.Info("welcome page requested")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome to fxEcho Example API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"health":     "/health",
			"users":      "/api/users",
			"user_by_id": "/api/users/:id",
		},
	})
}
