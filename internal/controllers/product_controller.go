package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
)

// SearchProduct godoc
// @Summary Search products
// @Description Search products by query (optional)
// @Tags products
// @Accept  json
// @Produce  json
// @Param query query string false "query string for searching products"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /products [get]
// PostgresController is a local type that wraps db.Postgres
func SearchProducts(e echo.Context) error {
	postgres := db.DB()
	query := e.QueryParam("query")
	if query == "" {
		products, err := postgres.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return e.JSON(http.StatusOK, products)

	}

	products, err := postgres.SearchProducts(query)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return e.JSON(http.StatusOK, products)
}
