package interfaces

import "github.com/netesh5/go_ecommerce/internal/models"

type IProduct interface {
	GetProductByID(id int) (models.Prouduct, error)
	GetAllProducts() ([]models.Prouduct, error)
}
