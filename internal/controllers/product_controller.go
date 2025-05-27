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
func (pc *PostgresController) SearchProducts(e echo.Context) error {

	query := e.QueryParam("query")
	if query == "" {
		products, err := pc.DB.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return e.JSON(http.StatusOK, products)

	}

	products, err := pc.DB.SearchProducts(query)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return e.JSON(http.StatusOK, products)
}

type PostgresController struct {
	DB *db.Postgres
}
