package utils

import (
	"fmt"
	"strings"

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

func ParseValidationError(err error) map[string]string {
	fieldErrors := make(map[string]string)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			field := strings.ToLower(fieldErr.Field())
			var msg string

			switch fieldErr.Tag() {
			case "required":
				msg = fmt.Sprintf("%s is required", fieldErr.Field())
			case "email":
				msg = "Invalid email format"
			case "min":
				msg = fmt.Sprintf("%s must be at least %s characters", fieldErr.Field(), fieldErr.Param())
			case "max":
				msg = fmt.Sprintf("%s must be at most %s characters", fieldErr.Field(), fieldErr.Param())
			default:
				msg = fmt.Sprintf("%s is invalid", fieldErr.Field())
			}

			fieldErrors[field] = msg
		}
	}

	return fieldErrors

}
