package models

type OTPData struct {
	Email string `json:"email" validate:"required"`
}

type VerfiyOTP struct {
	Email OTPData `json:"Email"  validate:"required"`
	Code  string  `json:"code" validate:"required"`
}
