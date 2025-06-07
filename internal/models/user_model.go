package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name" validate:"required,min=2,max=30"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone" validate:"required"`
	Password     string    `json:"password" validate:"required,min=6"`
	Verfiy       bool      `json:"is_verified"`
	Token        string    `json:"token"`
	Address      []Address `json:"address"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Address struct {
	Id        int    `json:"id"`
	Address   string `json:"address" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Country   string `json:"country" validate:"required"`
	ZipCode   string `json:"zip_code" validate:"required"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserLogin struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	Address      []Address `json:"address"`
	Token        string    `json:"token"`
	Verfiy       bool      `json:"is_verified"`
	RefreshToken string    `json:"refresh_token"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
