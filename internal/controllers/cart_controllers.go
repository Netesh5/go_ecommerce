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
// @Success 200 {object} responsehandler.SuccessResponse
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart/item [post]
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
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("product added to cart successfully"))

}

// RemoveItemFromCart godoc
// @Summary Remove product from cart
// @Description Remove a product from the user's cart
// @Tags cart
// @Accept  json
// @Produce  json
// @Param product_id path int true "Product ID"
// @Param user_id path int true "User ID"
// @Success 200 {object} responsehandler.SuccessResponse
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart/item [delete]
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
// @Success 200 {object} responsehandler.SuccessResponse{data=[]models.Cart} "List of Products in cart"
// @Failure 400 {object} errorhandler.ErrorHandler
// @Failure 500 {object} errorhandler.ErrorHandler
// @Router /cart [get]
func GetItemsFromCart(e echo.Context) error {
	postgres := db.DB()
	var cart models.Cart
	err := e.Bind(&cart)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}

	carts, err := postgres.GetItemsFromCart(cart.ID)
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
// @Success 200 {object} responsehandler.SuccessResponse
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

	carts, err := postgres.GetItemsFromCart(cart.ID)
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

// UpdateCartItem handles updating items in a user's cart.
//
// It binds the request to models.UpdateCartReq, validates required fields,
// and updates the cart item in the database for the authenticated user.
//
// @Summary Update an item in the user's cart
// @Description Updates quantity or other attributes of an item in the user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Param updateReq body models.UpdateCartReq true "Cart item update information"
// @Success 200 {object} responsehandler.SuccessResponse "Item updated successfully"
// @Failure 400 {object} errorhandler.ErrorHandler "Invalid input or required fields missing"
// @Failure 401 {object} errorhandler.ErrorHandler "Unauthorized"
// @Security BearerAuth
// @Router /cart/item [put]
func UpdateCartItem(e echo.Context) error {
	var updateReq models.UpdateCartReq

	if err := e.Bind(&updateReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid input"))
	}
	if err := e.Validate(&updateReq); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("required field are missing"))
	}

	user := e.Get("user").(models.User)
	db := db.DB()
	if err := db.UpdateCartItem(updateReq, user.ID); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("item update successfully"))
}
