package fxsupertoken

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// SupertokenMiddleware creates a middleware that checks if SuperTokens is initialized
func SupertokenMiddleware(config *SuperTokensConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !config.IsInitialized {
				// SuperTokens is not initialized, skip authentication
				c.Set("supertokensSession", nil)
				return next(c)
			}

			sess, err := session.GetSession(c.Request(), c.Response(), nil)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			}
			c.Set("supertokensSession", sess)
			return next(c)
		}
	}
}

// VerifySession creates a middleware that verifies SuperTokens session
func VerifySession(config *SuperTokensConfig) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !config.IsInitialized {
				// SuperTokens is not initialized, skip authentication
				c.Set("session", nil)
				return hf(c)
			}

			session.VerifySession(nil, func(rw http.ResponseWriter, r *http.Request) {
				c.Set("session", session.GetSessionFromRequestContext(r.Context()))

				// Call the handler
				err := hf(c)
				if err != nil {
					c.Error(err)
				}
			})(c.Response(), c.Request())

			return nil
		}
	}
}

// DefaultCORSHandler returns a custom CORS middleware with specific allowed origins
// Includes localhost ports for development and production UTOL domains
func DefaultCORSHandler() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			allowedOrigins := []string{
				"http://localhost:3000",
				"http://localhost:8000",
				"https://www.utol.com",
				"https://www.utol.com.ph",
				"https://portal-admin-v2.utol.com.ph",
				"https://staging-landing-page.utol.com.ph",
				"https://admin-staging-portal.utol.ph",
				"https://staging-admin-v2.utol.com.ph",
				"https://admin-portal.utol.com.ph",
				"https://staging-admin-portal.utol.com.ph",
				"https://staging-landing-v2.utol.com.ph",
				"https://accounting-admin.utol.com.ph",
				"https://staging-accounting-admin.utol.com.ph",
			}
			origin := c.Request().Header.Get("Origin")

			isAllowedOrigin := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					isAllowedOrigin = true
					break
				}
			}

			if isAllowedOrigin {
				c.Response().Header().Set("Access-Control-Allow-Origin", origin)
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if c.Request().Method == "OPTIONS" {
				c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...), ","))
				c.Response().Header().Set("Access-Control-Allow-Methods", "*")
				_, err := c.Response().Write([]byte(""))
				if err != nil {
					return err
				}
				return nil
			} else {
				return hf(c)
			}
		}
	}
}

// SuperTokensMiddlewareWrapper wraps SuperTokens middleware for Echo
func SuperTokensMiddlewareWrapper() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			supertokens.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				if err := hf(c); err != nil {
					c.Error(err)
				}
			})).ServeHTTP(c.Response(), c.Request())
			return nil
		}
	}
}
