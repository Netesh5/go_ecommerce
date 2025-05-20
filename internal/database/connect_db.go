package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/netesh5/go_ecommerce/internal/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbConfig.Host,
		cfg.DbConfig.Port,
		cfg.DbConfig.User,
		cfg.DbConfig.Password,
		cfg.DbConfig.DbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	log.Println("Connected to the database successfully")
	return db, nil
}
