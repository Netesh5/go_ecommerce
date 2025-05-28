package interfaces

import (
	"github.com/netesh5/go_ecommerce/internal/models"
)

type IAddress interface {
	UpdateAddress(models.Address) error
	DeleteAddress(id int) error
}
