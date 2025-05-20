package controllers

import "github.com/labstack/echo"

func GetProducts(e echo.Context) error {

	products := []string{"Product 1", "Product 2", "Product 3"}

	return e.JSON(200, products)

}
