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
}
