package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// AddReview godoc
// @Summary Add a review for a product
// @Description Creates a new review for a specific product
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param review body models.ReviewRequest true "Review Information"
// @Success 201 {object} responsehandler.SuccessResponse "Review added successfully"
// @Failure 400 {object} responsehandler.ErrorHandler "Invalid input, parameter or validation error"
// @Failure 500 {object} responsehandler.ErrorHandler "Server error"
// @Security BearerAuth
// @Router /products/{id}/reviews [post]
func AddReview(e echo.Context) error {
	var reviewReq models.ReviewRequest
	productIDParams := e.Param("id")
	if productIDParams == "" {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("product id is required"))
	}
	productID, err := strconv.Atoi(productIDParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid product id"))
	}

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
		ProductID: productID,
		Rating:    reviewReq.Rating,
		Comment:   reviewReq.Comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := db.AddReview(review); err != nil {
		return e.JSON(http.StatusInternalServerError, responsehandler.NewErrorHandler(err.Error()))
	}

	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("product review added successfully"))

}

// GetReviews godoc
// @Summary Get product reviews
// @Description Retrieves all reviews for a specific product by ID
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} responsehandler.SuccessResponse{data=[]models.Review} "Returns list of product reviews"
// @Failure 400 {object} responsehandler.ErrorHandler "Invalid request, missing or invalid product ID"
// @Failure 500 {object} responsehandler.ErrorHandler "Internal server error"
// @Router /products/{id}/reviews [get]
func GetReviews(e echo.Context) error {
	productIDParams := e.Param("id")
	if productIDParams == "" {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("product id is required"))
	}
	productID, err := strconv.Atoi(productIDParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid product id"))
	}
	db := db.DB()
	res, err := db.GetProductReviews(productID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(res, ""))
}
