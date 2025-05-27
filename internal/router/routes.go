package router

import (
	"net/http"

	"github.com/netesh5/go_ecommerce/internal/controllers"
)

var Routes = Routers{

	// {
	// 	Name:       "GetProducts",
	// 	Method:     http.MethodGet,
	// 	Path:       "/products",
	// 	HandleFunc: controllers.GetProducts,
	// },
	{
		Name:       "VerfyEmail",
		Method:     http.MethodPost,
		Path:       "/verify-email",
		HandleFunc: controllers.VerfiyEmail,
	},
	{
		Name:   "VerifyOTP",
		Method: http.MethodPost,
		Path:   "/verify-otp",
		// HandleFunc: controllers.VerifyOTP,
	},
	{
		Name:   "ResendEmail",
		Method: http.MethodPost,
		Path:   "/resend-email",
		// HandleFunc: controllers.ResendEmail,
	},
	{
		Name:       "SearchProducts",
		Method:     http.MethodGet,
		Path:       "/products",
		HandleFunc: controllers.SearchProducts,
	},
	// {
	// 	Name:       "GetProductByID",
	// 	Method:     http.MethodGet,
	// 	Path: 	 "/products/:id",
	// }
	{
		Name:       "AddProductToCart",
		Method:     http.MethodPost,
		Path:       "/cart",
		HandleFunc: controllers.AddProductToCart,
	},
	{
		Name:       "RemoveProductFromCart",
		Method:     http.MethodDelete,
		Path:       "/cart/:product_id",
		HandleFunc: controllers.RemoveProductFromCart,
	},
	{
		Name:       "GetCartItems",
		Method:     http.MethodGet,
		Path:       "/cart",
		HandleFunc: controllers.GetCartItems,
	},
}
