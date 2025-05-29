package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	contants "github.com/netesh5/go_ecommerce/internal/constant"
	"github.com/netesh5/go_ecommerce/internal/db"
	errorhandler "github.com/netesh5/go_ecommerce/internal/helper"
	"github.com/netesh5/go_ecommerce/internal/models"
	token "github.com/netesh5/go_ecommerce/internal/tokens"
	"golang.org/x/crypto/bcrypt"
)

// signup godoc
// @Summary singup
// @Description User signup
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserRequest true "User object"
// @Success 201 {object} map[string]interface{} "Returns success message"
// @Failure 400 {object} errorhandler.ErrorHandler "Validation error"
// @Failure 409 {object} errorhandler.ErrorHandler "User already exists"
// @Failure 500 {object} errorhandler.ErrorHandler "Internal server error"
// @Router /signup [post]
func SignUp(e echo.Context) error {
	var user models.User
	postgres := db.DB()
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
	res, err := postgres.GetUserByEmail(user.Email)
	if err == nil && res.ID != 0 {
		return e.JSON(http.StatusConflict, errorhandler.NewErrorHandler("user already exits"))
	}

	if err != nil && !errorhandler.IsNoRowsError(err) {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}

	password, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password

	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	token, refreshToken, _ := token.TokenGenerator(user.Email, user.Name, user.ID)
	user.Token = token
	user.RefreshToken = refreshToken
	// user.Cart = make([]models.Cart, 0)
	// user.Address = make([]models.Address, 0)
	// user.Orders = make([]models.Order, 0)

	_, err = postgres.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))

	}

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"sucess":  true,
		"message": "User created successfully",
	})
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Login godoc
// @Summary Login
// @Description User login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.UserLogin true "Login object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /login [post]
func Login(e echo.Context) error {
	var user models.UserLogin
	postgres := db.DB()
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, errorhandler.NewErrorHandler(
			err.Error(),
		))
	}
	res, err := postgres.GetUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(
			"user doesn't exists",
		))
	}

	passwordValid, msg := verfifyPassword(user.Password, res.Password)

	if passwordValid {
		return e.JSON(http.StatusInternalServerError, errorhandler.ErrorHandler{
			Message: msg.Error(),
		})
	}

	accessToken, refreshToken, err := token.TokenGenerator(res.Email, res.Name, res.ID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, errorhandler.NewErrorHandler(err.Error()))
	}

	res.Token = accessToken
	res.RefreshToken = refreshToken
	postgres.UpdateUser(res)

	userResponse := models.UserResponse{
		ID:           res.ID,
		Name:         res.Name,
		Email:        res.Email,
		Phone:        res.Phone,
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	return e.JSON(http.StatusOK, userResponse)

}

// VerfiyEmail godoc
// @Summary Verify Email
// @Description Verify email format
// @Tags auth
// @Accept  json
// @Produce  json
// @Param email query string true "Email address"
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
func verfifyPassword(userPassword string, currentPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(currentPassword)); err != nil {
		return false, fmt.Errorf("password is incorrect")
	}
	return true, nil
}
