package models

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}
