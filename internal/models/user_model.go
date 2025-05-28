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
	CreatedAt    time.Time `json:"created_at"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Address      []Address `json:"address"`
	Cart         []Cart    `json:"cart"`
	Orders       []Order   `json:"orders"`
}

type Address struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
