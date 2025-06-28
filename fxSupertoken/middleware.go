package fxsupertoken

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func SupertokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.GetSession(c.Request(), c.Response(), nil)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		}
		c.Set("supertokensSession", sess)
		return next(c)
	}
}

func VerifySession(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
