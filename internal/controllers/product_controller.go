package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// SearchProduct godoc
// @Summary Get Products / Search products
// @Description Search products by query (optional)
// @Tags products
// @Accept  json
// @Produce  json
// @Param query query string false "query string for searching products"
// @Success 200 {object} responsehandler.SuccessResponse{data=[]models.Product} "List of products"
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /products [get]
func SearchProducts(e echo.Context) error {
	postgres := db.DB()
	query := e.QueryParam("query")
	if query == "" {
		products, err := postgres.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
		}
		return e.JSON(http.StatusOK, responsehandler.SuccessWithData(products, ""))

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
// @Success 200 {object} responsehandler.SuccessResponse{data=models.Product} "Product"
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

// AddProduct godoc
// @Summary Add a new product
// @Description Creates a new product in the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.ProductReq true "Product information"
// @Success 201 {object} responsehandler.SuccessResponse "Product added successfully"
// @Failure 400 {object} responsehandler.ErrorHandler "Invalid input request or required fields are missing"
// @Failure 500 {object} responsehandler.ErrorHandler "Internal server error"
// @Router /products [post]
func AddProduct(e echo.Context) error {
	var productReq models.ProductReq

	if err := e.Bind(&productReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid input request"))
	}
	if err := e.Validate(&productReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("required fields are missing"))
	}
	product := models.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		Category:    productReq.Category,
		Stock:       productReq.Stock,
		Image:       productReq.Image,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	db := db.DB()

	if err := db.AddProduct(product); err != nil {
		return e.JSON(http.StatusInternalServerError, responsehandler.NewErrorHandler(err.Error()))
	}

	return e.JSON(http.StatusCreated, responsehandler.SuccessMessage("product added successfully"))

}

func AddReview(e echo.Context) error {
	var reviewReq models.ReviewRequest
	if err := e.Bind(&reviewReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid input request"))
	}
	if err := e.Validate(&reviewReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("required fields are missing"))
	}
	db := db.DB()

	user := e.Get("user").(models.User)

	review := models.Review{
		UserID:    user.ID,
		ProductID: reviewReq.ProductID,
		Rating:    reviewReq.Rating,
		Comment:   reviewReq.Comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
