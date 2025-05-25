package controllers

import (
	"github.com/labstack/echo/v4"
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
