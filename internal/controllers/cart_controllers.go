package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// AddItemToCart godoc
// @Summary Add product to cart
// @Description Add a product to the user's cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Param product_id path int true "Product ID"
// @Param user_id path int true "User ID"
// @Param cart body models.Cart true "Cart object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart [post]
func AddItemToCart(e echo.Context) error {
	postgres := db.DB()

	var cart models.Cart

	err := e.Bind(&cart)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}

	product, err := postgres.GetProductByID(cart.ProductID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	user, err := postgres.GetUserByID(cart.UserID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	if err := postgres.AddProductIntoCart(cart, product, user); err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Product added to cart successfully",
	})

}

// RemoveItemFromCart godoc
// @Summary Remove product from cart
// @Description Remove a product from the user's cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Param product_id path int true "Product ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart [delete]
func RemoveItemFromCart(e echo.Context) error {
	postgres := db.DB()
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

	if err := postgres.RemoveProductFromCart(productIdInt, userIdInt); err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("product removed from cart successfully"))
}

// GetItemFromCart godoc
// @Summary Get items from cart
// @Description Retrieve all items in the user's cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Cart
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart [get]
func GetItemFromCart(e echo.Context) error {
	postgres := db.DB()
	var cart models.Cart
	err := e.Bind(&cart)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}

	carts, err := postgres.GetItemFromCart(cart.ID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(carts, ""))
}

// BuyFromCart godoc
// @Summary Get items from cart
// @Description Retrieve all items in the user's cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Param user_id query int true "User ID"
// @Success 200 {array} models.Cart
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /buy-cart [get]
func BuyFromCart(e echo.Context) error {
	postgres := db.DB()

	var cart models.Cart
	err := e.Bind(&cart)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}

	carts, err := postgres.GetItemFromCart(cart.ID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}

	for _, cart := range carts {
		if err := postgres.RemoveProductFromCart(cart.ProductID, cart.UserID); err != nil {
			return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
		}

	}
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("products purchased successfully"))
}
