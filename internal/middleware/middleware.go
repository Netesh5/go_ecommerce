package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	token "github.com/netesh5/go_ecommerce/internal/tokens"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			clientToken := e.Request().Header.Get("token")
			if clientToken == "" {
				return e.JSON(http.StatusUnauthorized, errorhandler.NewErrorHandler("No authorization header found"))
			}

			claims, err := token.TokenValidator(clientToken)
			if err != "" {
				return e.JSON(http.StatusUnauthorized, errorhandler.NewErrorHandler(err))
			}
			e.Set("email", claims.Email)
			e.Set("id", claims.ID)
			return next(e)
		}
	}
}
