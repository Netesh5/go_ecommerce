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

// DeleteUserAddress godoc
// @Summary Delete a user's address
// @Description Removes a specified address associated with a user
// @Tags addresses
// @Accept json
// @Produce json
// @Param id path int true "Address information containing address Id"
// @Success 200 {object} responsehandler.SuccessResponse
// @Failure 400 "Invalid input or deletion error"
// @Router /address/{id} [delete]
func DeleteUserAddress(e echo.Context) error {
	postgres := db.DB()
	var address models.Address
	err := e.Bind(&address)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}
	if err := postgres.DeleteAddress(address.Id, address.UserId); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("address deleted successfully"))
}

// GetAddress godoc
// @Summary Get a user's address
// @Description Retrieves a specific address associated with the authenticated user
// @Tags addresses
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} responsehandler.SuccessResponse{data=models.Address} "User Address"
// @Failure 400 {object} errorhandler.ErrorHandler "Bad request"
// @Failure 404 {string} string "No data found"
// @Router /address/{id} [get]
func GetAddressByID(e echo.Context) error {
	postgres := db.DB()

	// Get address ID from path
	addressParam := e.Param("id")
	if addressParam == "" {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler("address id is required"))
	}
	addressId, err := strconv.Atoi(addressParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "invalid address id")
	}

	// Get authenticated user from context
	user := e.Get("user").(models.User)

	address, err := postgres.GetUserAddress(addressId, user.ID)
	if err != nil {
		return e.JSON(http.StatusNotFound, "no data found")
	}

	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(address, ""))
}

// GetAddresses godoc
// @Summary Get user addresses
// @Description Retrieve all addresses for the authenticated user
// @Tags addresses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responsehandler.SuccessResponse{data=[]models.Address} "User Address"
// @Failure 500 {string} string "Error message"
// @Router /addresses [get]
func GetAddresses(e echo.Context) error {
	userId := e.Get("user").(models.User)

	db := db.DB()

	addresses, err := db.GetUserAddresses(userId.ID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, responsehandler.SuccessWithData(addresses, ""))

}

// AddAddress godoc
// @Summary Add a new address for the authenticated user
// @Description Adds a new address to the user's profile
// @Tags address
// @Accept json
// @Produce json
// @Param address body models.Address true "Address information"
// @Success 200 {object} responsehandler.SuccessResponse "Success message with address updated confirmation"
// @Failure 400 {object} errorhandler.ErrorHandler "Invalid address input"
// @Failure 500 {object} errorhandler.ErrorHandler "Internal server error"
// @Security ApiKeyAuth
// @Router /address [post]
func AddAddress(e echo.Context) error {
	var address models.Address

	if err := e.Bind(&address); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid address input"))
	}
	if err := e.Validate(&address); err != nil {
		return e.JSON(http.StatusBadRequest, responsehandler.NewErrorHandler("invalid address input"))
	}

	db := db.DB()

	user := e.Get("user").(models.User)

	address.UserId = user.ID

	if err := db.AddUserAddress(address); err != nil {
		return e.JSON(http.StatusInternalServerError, responsehandler.NewErrorHandler(err.Error()))
	}

	return e.JSON(http.StatusOK, responsehandler.SuccessMessage("address updated successfully"))
}
