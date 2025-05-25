package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

func AddToCart(e echo.Context, db db.Postgres, cart models.Cart) error {
	productId := e.Param("id")
	userId := e.Param("productId")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{
			Message: "Invalid product ID",
		})
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{
			Message: "Invalid user ID",
		})
	}

	product, err := db.GetProductByID(productIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}

	user, err := db.GetUserByID(userIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}

	if err := db.AddProductIntoCart(cart, product, user); err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Product added to cart successfully",
	})

}

func RemoveItem() error {

}

func GetItemFromCart() error {

}

func BuyFromCart() error {

}
