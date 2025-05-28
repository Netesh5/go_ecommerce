package controllers

import (
	"net/http"

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
// @Param address body models.Address true "Address information containing Id and UserId"
// @Success 200 "Address successfully deleted"
// @Failure 400 {object} errorhandler.ErrorResponse "Invalid input or deletion error"
// @Router /user/address [delete]
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
