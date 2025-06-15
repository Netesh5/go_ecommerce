package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/services"
)

func UploadImage(e echo.Context) error {
	file, fileHeader, err := e.Request().FormFile("image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}

	defer file.Close()

	imgURL, err := services.UploadImageToCloudinary(file, fileHeader)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, responsehandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData("", imgURL))
}
