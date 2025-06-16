package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
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
	product, err := db.GetProductByID(productId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}

}
