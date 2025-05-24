package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userdb "github.com/netesh5/go_ecommerce/internal/database"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

func SignUp(e echo.Context, db userdb.Postgres) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}

	if err := e.Validate(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}
	res, err := db.GetUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})
	}
	if res.ID != 0 {
		return e.JSON(http.StatusConflict, errorhandler.ErrorHandler{
			Message: "User already exists",
		})
	}

}
