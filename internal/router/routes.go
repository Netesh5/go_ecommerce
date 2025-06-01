package router

import (
	"net/http"

	"github.com/netesh5/go_ecommerce/internal/controllers"
)

var Routes = Routers{

	{
		Name:       "SignUp",
		Method:     http.MethodPost,
		Path:       "/signup",
		HandleFunc: controllers.SignUp,
	},
	{
		Name:       "Login",
		Method:     http.MethodPost,
		Path:       "/login",
		HandleFunc: controllers.Login,
	},
	// {
	// 	Name:       "VerfyEmail",
	// 	Method:     http.MethodPost,
	// 	Path:       "/verify-email",
	// 	HandleFunc: controllers.VerfiyEmail,
	// },
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
		HandleFunc: controllers.AddItemToCart,
	},
	{
		Name:       "RemoveProductFromCart",
		Method:     http.MethodDelete,
		Path:       "/cart",
		HandleFunc: controllers.RemoveItemFromCart,
	},
	{
		Name:       "GetCartItems",
		Method:     http.MethodGet,
		Path:       "/cart",
		HandleFunc: controllers.GetItemFromCart,
	},

	{
		Name:       "DeleteUserAddress",
		Method:     http.MethodDelete,
		Path:       "/address/:id",
		HandleFunc: controllers.DeleteUserAddress,
	},
	{
		Name:       "GetUserAddressById",
		Method:     http.MethodGet,
		Path:       "/address/:id",
		HandleFunc: controllers.GetAddressByID,
	},
	{
		Name:       "GetUserAddresses",
		Method:     http.MethodGet,
		Path:       "/addresses",
		HandleFunc: controllers.GetAddresses,
	},
}
