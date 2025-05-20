package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/netesh5/go_ecommerce/internal/config"
	"github.com/netesh5/go_ecommerce/internal/router"
)

func main() {
	config := config.LoadConfig()
	e := echo.New()

	router.RegisterRoutes(e, router.Routes, config.ApiVersion)
	log.Println("Server is running on port", config.Server.Address)
	e.Logger.Fatal(e.Start(config.Server.Address))

}
