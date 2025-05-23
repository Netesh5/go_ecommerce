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

func insertData(db *sql.DB, InsertModel interface{}) (*sql.DB, error) {
	// Example of inserting data into a table
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := db.Exec(query, "John Doe")
	if err != nil {
		log.Fatal("Error inserting data:", err)
		return nil, err
	}
	log.Println("Data inserted successfully")
	return db, nil

}
