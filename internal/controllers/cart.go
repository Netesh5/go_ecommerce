package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
)

func AddToCart(e echo.Context) error {
	productId := e.Param("id")
	if productId == "" {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{
			Message: "Product ID is required",
		})
	}

}

func RemoveItem() error {

}

func GetItemFromCart() error {

}

func BuyFromCart() error {

}
