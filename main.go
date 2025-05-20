package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/netesh5/go_ecommerce/internal/config"
	"github.com/netesh5/go_ecommerce/internal/database"
	"github.com/netesh5/go_ecommerce/internal/router"
)

func main() {
	config := config.LoadConfig()
	e := echo.New()

	router.RegisterRoutes(e, router.Routes, config.ApiVersion)
	log.Println("Server is running on port", config.Server.Address)
	e.Logger.Fatal(e.Start(config.Server.Address))
	db, err := database.ConnectDB(config)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	log.Println("Database connection closed")
	log.Println("Server stopped")
}
