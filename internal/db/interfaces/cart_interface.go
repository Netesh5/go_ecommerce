package interfaces

import "github.com/netesh5/go_ecommerce/internal/models"

type ICart interface {
	AddProductIntoCart(models.Prouduct) error
	RemoveProductFromCart(productID int, userID int) error
}
