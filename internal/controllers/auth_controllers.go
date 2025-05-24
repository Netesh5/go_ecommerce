package controllers

import (
	"net/http"
	"time"

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

	password := HashPassword(user.Password)
	user.Password = password

	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	token, refreshToken, _ := generate.TokenGenerator(user.Email, user.Name, user.ID)
	user.Token = token
	user.RefreshToken = refreshToken
	user.Cart = make([]models.Cart, 0)
	user.Address = models.Address{}
	user.Orders = make([]models.Order, 0)

	newUser, err := db.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})

	}
	return e.JSON(http.StatusCreated, newUser)
}

func HashPassword(password *string) (*string, error) {
	// Implement password hashing logic here
	// For example, using bcrypt:
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	//     return "", err
	// }
	// return string(hashedPassword), nil

	// Placeholder return for now
	return password, nil
}
