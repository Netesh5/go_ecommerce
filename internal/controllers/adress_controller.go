package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
)

func DeleteAddress(e echo.Context) error {
	postgres := db.DB()
	paramId := e.Param("id")
	if paramId == "" {
		return e.JSON(http.StatusNotFound, errorhandler.NewErrorHandler("Id must be provided"))
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler("Invalid request ID"))
	}
	if err := postgres.DeleteAddress(id); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(err.Error()))
	}
	return nil
}
