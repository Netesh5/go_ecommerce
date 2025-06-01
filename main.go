package main

import (
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/netesh5/go_ecommerce/docs"
	"github.com/netesh5/go_ecommerce/internal/config"
	userdb "github.com/netesh5/go_ecommerce/internal/db"
	"github.com/netesh5/go_ecommerce/internal/middleware"
	"github.com/netesh5/go_ecommerce/internal/router"
	"github.com/netesh5/go_ecommerce/internal/services"
	"github.com/netesh5/go_ecommerce/internal/utils"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Ecommerce API Swagger
// @version 1.0
// @description This is a sample server using Echo.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config := config.LoadConfig()
	e := echo.New()
	e.Validator = utils.NewValidator()
	e.HTTPErrorHandler = middleware.ErrorHandler

	e.Use(middleware.LogrusLogger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	db, err := userdb.ConnectDB(config)

	services.CreateTwiiloClient(config.OtpConfig.AccountSID, config.OtpConfig.AuthToken, config.OtpConfig.ServiceSID)

	router.RegisterRoutes(e, router.Routes, config.ApiVersion)
	logrus.Info("Server is running on port", config.Server.Address)
	e.Logger.Fatal(e.Start(config.Server.Address))

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	defer db.Db.Close()
	log.Println("Database connection closed")
	log.Println("Server stopped")
}
