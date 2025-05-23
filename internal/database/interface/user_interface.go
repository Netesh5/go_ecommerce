package user_interface

import "github.com/netesh5/go_ecommerce/internal/types"

type UserInterface interface {
	GetUserByEmail(email string) (types.User, error)
	GetUserByID(id int) (types.User, error)
	CreateUser(user types.User) (types.User, error)
	UpdateUser(user types.User) (types.User, error)
	DeleteUser(id int) error
}
