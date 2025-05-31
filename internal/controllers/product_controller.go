package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
)

// SearchProduct godoc
// @Summary Get Products / Search products
// @Description Search products by query (optional)
// @Tags products
// @Accept  json
// @Produce  json
// @Param query query string false "query string for searching products"
// @Success 200 {array} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [get]
// PostgresController is a local type that wraps db.Postgres
func SearchProducts(e echo.Context) error {
	postgres := db.DB()
	query := e.QueryParam("query")
	if query == "" {
		products, err := postgres.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
		}
		return e.JSON(http.StatusOK, products)

	}

	products, err := postgres.SearchProducts(query)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(products, ""))
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Retrieves a product from the database based on the provided ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad request - Invalid or missing product ID"
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /products/{id} [get]
func GetProductByID(e echo.Context) error {
	productParam := e.Param("id")
	if productParam == "" {
		return e.JSON(http.StatusBadRequest, "product id is required")
	}
	productId, err := strconv.Atoi(productParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "product id is invalid")
	}

	db := db.DB()
	product, err := db.GetProductByID(productId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(product, ""))
}
