package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
)

// // getUser godoc
// // @Summary Get a user
// // @Description Get user by ID
// // @Tags users
// // @Accept  json
// // @Produce  json
// // @Param id path int true "User ID"
// // @Success 200 {object} map[string]interface{}
// // @Failure 400 {object} map[string]string
// // @Router /user/{id} [get]
// func GetProducts(e echo.Context) error {

// 	products := []string{"Product 1", "Product 2", "Product 3"}

// 	return e.JSON(200, products)

// }

// SearchProduct godoc
// @Summary Search products
// @Description Search products by query
// @Tags products
// @Accept  json
// @Produce  json
// @Param query query string true "query string for searching products"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /products [get]
func SearchProducts(e echo.Context, db *db.Postgres) error {

	query := e.QueryParam("query")
	if query == "" {
		products, err := db.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return e.JSON(http.StatusOK, products)

	}

	products, err := db.SearchProducts(query)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return e.JSON(http.StatusOK, products)
}
