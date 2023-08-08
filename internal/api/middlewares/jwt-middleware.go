package middlewares

import (
	"net/http"

	jwttoken "github.com/andrersp/go-stock/internal/pkg/jwt-token"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user, err := jwttoken.VerifyJwtToken(c.Request())

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		c.Set("token", *user)

		return next(c)
	}
}
