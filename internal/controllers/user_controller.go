package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	responsehandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// GetUser godoc
// @Summary Get user information
// @Description Retrieves the current user's information from the database
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responsehandler.SuccessResponse{data=models.User} "User data successfully retrieved"
// @Failure 400 {object} responsehandler.ErrorHandler "Error retrieving user data"
// @Router /user/getme [get]
func GetUser(e echo.Context) error {
	user := e.Get("user").(models.User)
	db := db.DB()
	res, err := db.GetUserByID(user.ID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(res, ""))
}

// @Summary Update a user's information
// @Description Updates user information in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserUpdate true "User information"
// @Success 200 {object} responsehandler.SuccessResponse "User info updated successfully"
// @Failure 400 {object} responsehandler.ErrorHandler "Invalid input / Required fields are missing / Couldn't update user info"
// @Router /user [put]
func UpdateUser(e echo.Context) error {
	var user models.UserUpdate

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid input"))
	}

	if err := e.Validate(&user); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("required fields are missing"))
	}

	db := db.DB()

	_, err := db.UpdateUserInfo(user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("couldn't update user info"))
	}
	return e.JSON(http.StatusBadRequest, responsehandler.SuccessMessage("user info updated successfully"))
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Allows authenticated users to update their password by providing current password and new password
// @Tags users
// @Accept json
// @Produce json
// @Param updatePassword body models.UpdatePassword true "Password update information"
// @Success 200 {object} responsehandler.SuccessResponse "Password updated successfully"
// @Failure 400 {object} responsehandler.ErrorHandler "Error messages for invalid input, missing fields, password mismatch, user not found, incorrect current password, or update failure"
// @Security BearerAuth
// @Router /users/password [put]
func UpdatePassword(e echo.Context) error {
	var updatePassword models.UpdatePassword

	if err := e.Bind(&updatePassword); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid input"))
	}
	if err := e.Validate(&updatePassword); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("required fields are missing"))
	}

	isMatch := updatePassword.NewPassword == updatePassword.ConfirmPassword
	if !isMatch {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("new password and confirm password doesn't match"))
	}
	user := e.Get("user").(models.User)

	db := db.DB()

	user, err := db.GetUserByID(user.ID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("user not found"))
	}

	_, err = verfifyPassword(updatePassword.CurrentPassword, user.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("incorrect current password"))
	}

	hashPass, err := HashPassword(updatePassword.NewPassword)
	if err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("couldn't update the password"))
	}
	if err := db.UpdatePassword(user.Email, hashPass); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("couldn't update the password"))
	}

	return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("password updated successfully"))

}
