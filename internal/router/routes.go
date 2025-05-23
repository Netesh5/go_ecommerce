package router

import (
	"net/http"

	"github.com/netesh5/go_ecommerce/internal/controllers"
)

var Routes = Routers{

	{
		Name:       "GetProducts",
		Method:     http.MethodGet,
		Path:       "/products",
		HandleFunc: controllers.GetProducts,
	},
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
}
