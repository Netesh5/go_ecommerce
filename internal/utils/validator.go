package utils

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

// For creating a instnce of the struct
func NewValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// For using in controller to validate the struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
