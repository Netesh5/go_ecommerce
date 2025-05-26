package controllers

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/labstack/echo/v4"
	contants "github.com/netesh5/go_ecommerce/internal/constant"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
)

// signup godoc
// @Summary singup
// @Description User signup
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /signup [post]

func SignUp(e echo.Context, db db.Postgres) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	if err := e.Validate(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	res, err := db.GetUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	if res.ID != 0 {
		return e.JSON(http.StatusConflict, errorhandler.ErrorHandler{
			Message: "User already exists",
		})
	}

	// password, err := HashPassword(user.Password)
	// user.Password = password

	// user.CreatedAt = time.Now().UTC()
	// user.UpdatedAt = time.Now().UTC()

	// token, refreshToken, _ := generate.TokenGenerator(user.Email, user.Name, user.ID)
	// user.Token = token
	// user.RefreshToken = refreshToken
	user.Cart = make([]models.Cart, 0)
	user.Address = models.Address{}
	user.Orders = make([]models.Order, 0)

	_, err = db.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: err.Error(),
		})

	}

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"sucess":  true,
		"message": "User created successfully",
	})
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

// Login godoc
// @Summary Login
// @Description User login
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /login [post]

func login(e echo.Context, db db.Postgres) error {
	var user models.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	if err := e.Validate(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	res, err := db.GetUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			"user doesn't exists",
		))
	}

	// passwordValid, msg := verfifyPassword(user.Password, res.Password)

	// if passwordValid {
	// 	return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
	// 		Message: msg,
	// 	})
	// }

	// token, refresh := generate.TokenGenerator(user.Email, user.Name, user.ID)

	// generate.UpdateAllToken(token, refreshToken, res.ID)

	return e.JSON(http.StatusOK, res)

}

// VerfiyEmail godoc
// @Summary Verify Email
// @Description Verify email format
// @Tags users
// @Accept  json
// @Produce  json
// @Param email path string true "Email address"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /verify-email [get]

func VerfiyEmail(e echo.Context) error {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emailParam := e.Param("email")
	email, err := url.QueryUnescape(emailParam)
	if err != nil || email == "" {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(contants.EmailValidaionError))
	}

	// Validate email format
	if !emailRegex.MatchString(email) {
		e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(contants.EmailValidaionError))
	}

	return nil
}
