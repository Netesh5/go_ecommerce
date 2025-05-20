package main

import (
	"github.com/labstack/echo"
	"github.com/netesh5/go_ecommerce/internal/config"
	"github.com/netesh5/go_ecommerce/internal/router"
)

func main() {
	config := config.LoadConfig()
	println("Config loaded successfully", config.Env)
	e := echo.New()

	routes := router.Routes{
		{
			Name:   "GetProducts",
			Method: "GET",
			Path:   "/products",
			HandleFunc: func(c echo.Context) error {
				return c.String(200, "Hello, World!")
			},
		},
	}
	router.RegisterRoutes(e, routes, config.ApiVersion)
}
