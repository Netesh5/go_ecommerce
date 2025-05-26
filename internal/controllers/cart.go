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
	productId := e.Param("product_id")
	userId := e.Param("user_id")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(

			"Invalid product ID",
		))
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			"Invalid user ID",
		))
	}

	product, err := db.GetProductByID(productIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	user, err := db.GetUserByID(userIdInt)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	if err := db.AddProductIntoCart(cart, product, user); err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Product added to cart successfully",
	})

}

func RemoveItem(e echo.Context, db *db.Postgres) error {
	productId := e.Param("product_id")
	userId := e.Param("user_id")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			"Invalid product ID",
		))
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			"Invalid user ID",
		))
	}

	if err := db.RemoveProductFromCart(productIdInt, userIdInt); err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Product removed from cart successfully",
	})
}

func GetItemFromCart(e echo.Context, db *db.Postgres) error {
	userId := e.Param("user_id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler("Invalid user ID"))
	}
	carts, err := db.GetItemFromCart(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, carts)
}

func BuyFromCart() error {

}
