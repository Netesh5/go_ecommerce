package controllers

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/labstack/echo/v4"
	contants "github.com/netesh5/go_ecommerce/internal/constant"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
)

// getUser godoc
// @Summary Get a user
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /user/{id} [get]
func GetProducts(e echo.Context) error {

	products := []string{"Product 1", "Product 2", "Product 3"}

	return e.JSON(200, products)

}

func VerfiyEmail(e echo.Context) error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emailParam := e.Param("email")
	email, err := url.QueryUnescape(emailParam)
	if err != nil || email == "" {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{Error: true, Message: contants.EmailValidaionError})
	}

	// Validate email format
	if !emailRegex.MatchString(email) {
		e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{Error: true, Message: contants.EmailValidaionError})
	}

	return nil
}
