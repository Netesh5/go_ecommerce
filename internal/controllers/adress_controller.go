package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// DeleteUserAddress godoc
// @Summary Delete a user's address
// @Description Removes a specified address associated with a user
// @Tags addresses
// @Accept json
// @Produce json
// @Param address path id true "Address information containing address Id"
// @Success 200 "Address successfully deleted"
// @Failure 400 "Invalid input or deletion error"
// @Router /user/{id} [delete]
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
	return nil
}

// GetAddress godoc
// @Summary Get a user's address
// @Description Retrieves a specific address associated with the authenticated user
// @Tags addresses
// @Accept json
// @Produce json
// @Param id path string true "Address ID"
// @Success 200 {object} models.Address "Address details"
// @Failure 400 {object} errorhandler.ErrorResponse "Bad request"
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

	return e.JSON(http.StatusOK, address)
}
