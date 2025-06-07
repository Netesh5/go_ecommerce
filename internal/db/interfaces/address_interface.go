package interfaces

import (
	"github.com/netesh5/go_ecommerce/internal/models"
)

type IAddress interface {
	UpdateUserAddress(models.Address) error
	DeleteUserAddress(id int) error
	GetUserAddress(id int) (models.Address, error)
	GetUserAddresses(userId int) ([]models.Address, error)
	AddUserAddress(models.Address) error
}
