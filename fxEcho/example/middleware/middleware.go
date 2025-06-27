package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// NewRequestTimingMiddleware creates middleware for request timing
func NewRequestTimingMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)

			logger.Info("request processed",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Int("status", c.Response().Status),
				zap.Duration("duration", duration),
			)

			return err
		}
	}
}

// NewRequestIDMiddleware creates middleware for adding request ID
func NewRequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = time.Now().Format("20060102150405") + "-" + c.RealIP()
			}
			c.Response().Header().Set("X-Request-ID", requestID)
			return next(c)
		}
	}
}
