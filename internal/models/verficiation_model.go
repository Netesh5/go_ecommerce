package models

type Verfication struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	OTP      string `json:"otp"`
	Verified bool   `json:"verified"`
}
