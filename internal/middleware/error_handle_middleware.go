package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/utils"
)

func ErrorHandler(err error, c echo.Context) {
	// Validation error
	if validationErrors := utils.ParseValidationError(err); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": "Validation failed",
			"fields":  validationErrors,
		})
		return
	}

	// Custom auth/functional error (set as HTTPError)
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, echo.Map{
			"error":   true,
			"message": he.Message,
		})
		return
	}

	// Fallback: internal server error
	c.JSON(http.StatusInternalServerError, echo.Map{
		"error":   true,
		"message": "Internal Server Error",
	})
}
