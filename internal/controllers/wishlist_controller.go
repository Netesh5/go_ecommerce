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
