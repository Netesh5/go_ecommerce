package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
	"github.com/netesh5/go_ecommerce/internal/services"
)

// UploadImage handles the file upload process to Cloudinary
//
// @Summary Upload an image to Cloudinary
// @Description Receives an image file from form data and uploads it to Cloudinary
// @Tags images
// @Accept mpfd
// @Produce json
// @Param image formData file true "Image file to upload"
// @Success 200 {object} models.ImageResponse "Returns the URL of the uploaded image"
// @Failure 400 {object} responsehandler.ErrorHandler "Invalid request or missing file"
// @Failure 500 {object} responsehandler.ErrorHandler "Server error during upload"
// @Router /put-image [post]
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
	imageRes := models.ImageResponse{
		ImageUrl: imgURL,
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(imageRes, ""))
}
