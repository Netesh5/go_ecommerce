package interfaces

import (
	"github.com/netesh5/go_ecommerce/internal/models"
)

type IAddress interface {
	UpdateUserAddress(models.Address) error
	DeleteUserAddress(id int) error
}
