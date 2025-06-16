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

// AddProductToWishList godoc
// @Summary Add a product to the user's wishlist
// @Description Add a product to the authenticated user's wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Success 200 {object} responsehandler.SuccessResponse "Product added to wishlist"
// @Failure 400 {object} responsehandler.ErrorHandler "Bad request"
// @Router /wishlist/{id} [post]
func AddProductToWishList(e echo.Context) error {
	productParam := e.Param("id")
	if productParam == "" {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("product id is required"))
	}

	productId, err := strconv.Atoi(productParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid product id"))
	}

	db := db.DB()
	_, err = db.GetProductByID(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}
	userID := e.Get("user").(models.User)

	wishlist := models.Wishlists{
		UserId:    userID.ID,
		ProductId: productId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if err := db.AddProductToWishList(wishlist); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("product added to wishlist"))
}

// GetUserWishlist godoc
// @Summary Get user's wishlist products
// @Description Retrieves all products in the authenticated user's wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responsehandler.SuccessResponse{data=[]models.Product} "Successfully retrieved wishlist products"
// @Failure 400 {object} responsehandler.ErrorHandler "Error retrieving wishlist"
// @Router /wishlist [get]
func GetUserWishlist(e echo.Context) error {
	userID := e.Get("user").(models.User)

	db := db.DB()

	wishlists, err := db.GetUserWishlistProducts(userID.ID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}

	var products []models.Product

	for _, wishlist := range wishlists {
		product, err := db.GetProductByID(wishlist.ProductId)
		if err != nil {
			continue
		}
		products = append(products, product)
	}

	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(products, ""))
}
