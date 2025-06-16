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
	{
		Name:       "VerfyOTP",
		Method:     http.MethodPost,
		Path:       "/auth/verify-email-otp",
		HandleFunc: controllers.VerifyEmailVerificationOTP,
	},
	{
		Name:       "SendOTP",
		Method:     http.MethodPost,
		Path:       "/auth/send-email-otp",
		HandleFunc: controllers.SendEmailVerificationOTP,
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

	{
		Name:       "AddProductToCart",
		Method:     http.MethodPost,
		Path:       "/cart/item",
		HandleFunc: controllers.AddItemToCart,
	},
	{
		Name:       "RemoveProductFromCart",
		Method:     http.MethodDelete,
		Path:       "/cart/item",
		HandleFunc: controllers.RemoveItemFromCart,
	},
	{
		Name:       "GetCartItems",
		Method:     http.MethodGet,
		Path:       "/cart",
		HandleFunc: controllers.GetItemsFromCart,
	},
	{
		Name:       "UpateCartItem",
		Method:     http.MethodPut,
		Path:       "/cart/item",
		HandleFunc: controllers.UpdateCartItem,
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
	{
		Name:       "ForgetPassword",
		Method:     http.MethodPost,
		Path:       "/auth/forget-password",
		HandleFunc: controllers.ForgetPassword,
	},
	{
		Name:       "VerfiyPasswordOTP",
		Method:     http.MethodPost,
		Path:       "/auth/verify-reset-otp",
		HandleFunc: controllers.VerifyPasswordResetOtp,
	},
	{
		Name:       "ResetPassword",
		Method:     http.MethodPost,
		Path:       "/auth/reset-password",
		HandleFunc: controllers.ResetPassword,
	},
	{
		Name:       "AddUserAddress",
		Method:     http.MethodPost,
		Path:       "/address",
		HandleFunc: controllers.AddAddress,
	},
	{
		Name:       "GetUser",
		Method:     http.MethodGet,
		Path:       "/user/get-me",
		HandleFunc: controllers.GetUser,
	},
	{
		Name:       "UpdateUser",
		Method:     http.MethodPut,
		Path:       "/user",
		HandleFunc: controllers.UpdateUser,
	},
	{
		Name:       "UpdatePassword",
		Method:     http.MethodPut,
		Path:       "/user/password",
		HandleFunc: controllers.UpdatePassword,
	},
	{
		Name:       "AddProduct",
		Method:     http.MethodPost,
		Path:       "/products",
		HandleFunc: controllers.AddProduct,
	},
	{
		Name:       "PutImage",
		Method:     http.MethodPost,
		Path:       "/put-image",
		HandleFunc: controllers.UploadImage,
	},
	{
		Name:       "AddReview",
		Method:     http.MethodPost,
		Path:       "/products/:id/review",
		HandleFunc: controllers.AddReview,
	},
}
