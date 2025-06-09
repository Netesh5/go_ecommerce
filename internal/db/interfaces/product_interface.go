package interfaces

import "github.com/netesh5/go_ecommerce/internal/models"

type IProduct interface {
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProducts(models.Product) error
}
