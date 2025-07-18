package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
	token "github.com/netesh5/go_ecommerce/internal/tokens"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			authHeader := e.Request().Header.Get("Authorization")
			clientToken := ""
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				clientToken = authHeader[7:]
			}
			if clientToken == "" {
				return e.JSON(http.StatusUnauthorized, errorhandler.NewErrorHandler("No authorization header found"))
			}

			claims, err := token.TokenValidator(clientToken)
			if err != "" {
				return e.JSON(http.StatusUnauthorized, errorhandler.NewErrorHandler(err))
			}
			// e.Set("email", claims.Email)
			// e.Set("id", claims.ID)
			// e.Set("name", claims.Name)

			user := models.User{
				ID:    claims.ID,
				Name:  claims.Name,
				Email: claims.Email,
			}
			e.Set("user", user)
			return next(e)
		}
	}
}
