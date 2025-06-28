// Health check HTTP handlers for the API.
package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	db     *gorm.DB
	logger *zap.Logger
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *gorm.DB, logger *zap.Logger) *HealthHandler {
	return &HealthHandler{
		db:     db,
		logger: logger,
	}
}

// HealthCheck handles health check requests
func (h *HealthHandler) HealthCheck(c echo.Context) error {
	status := "healthy"
	httpStatus := http.StatusOK

	// Check database connection
	dbStatus := "healthy"
	if err := h.db.Raw("SELECT 1").Error; err != nil {
		dbStatus = "unhealthy"
		status = "unhealthy"
		httpStatus = http.StatusServiceUnavailable
		h.logger.Error("database health check failed", zap.Error(err))
	}

	response := map[string]interface{}{
		"status":    status,
		"timestamp": time.Now().UTC(),
		"services": map[string]interface{}{
			"database": dbStatus,
		},
		"version": "1.0.0",
	}

	return c.JSON(httpStatus, response)
}

// ReadinessCheck handles readiness check requests
func (h *HealthHandler) ReadinessCheck(c echo.Context) error {
	status := "ready"
	httpStatus := http.StatusOK

	// Check if database is ready
	if err := h.db.Raw("SELECT 1").Error; err != nil {
		status = "not ready"
		httpStatus = http.StatusServiceUnavailable
		h.logger.Error("readiness check failed", zap.Error(err))
	}

	response := map[string]interface{}{
		"status":    status,
		"timestamp": time.Now().UTC(),
	}

	return c.JSON(httpStatus, response)
}
