package models

type ForgetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type ResetPassword struct {
	Email           ForgetPasswordRequest `json:"email" validate:"required"`
	NewPassword     string                `json:"password" validate:"required"`
	ConfirmPassword string                `json:"confirm_password" validate:"required"`
}

type UpdatePassword struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
