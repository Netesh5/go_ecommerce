package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/netesh5/go_ecommerce/internal/config"
)

type Postgres struct {
	Db *sql.DB
}

var database *Postgres

func ConnectDB(cfg *config.Config) (*Postgres, error) {
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
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT email_format CHECK (email ~* '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$')

	)`)
	log.Println("Connected to the database successfully")
	return &Postgres{
		Db: db,
	}, nil

}

func DB() *Postgres {
	return &Postgres{
		Db: database.Db,
	}

}

// func insertData(db *sql.DB, InsertModel interface{}) (*sql.DB, error) {
// 	// Example of inserting data into a table
// 	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
// 	_, err := db.Exec(query, "John Doe")
// 	if err != nil {
// 		log.Fatal("Error inserting data:", err)
// 		return nil, err
// 	}
// 	log.Println("Data inserted successfully")
// 	return db, nil

// }
