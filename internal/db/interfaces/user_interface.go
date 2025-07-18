package interfaces

import "github.com/netesh5/go_ecommerce/internal/models"

type IUser interface {
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}
