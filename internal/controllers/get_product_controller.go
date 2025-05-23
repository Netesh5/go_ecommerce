package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	contants "github.com/netesh5/go_ecommerce/internal/constant"
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
	email := e.Param("email")
	if email == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": contants.EmailValidaionError})
	}

}
