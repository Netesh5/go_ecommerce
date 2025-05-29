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
	_, dbErr := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		phone VARCHAR(15) NOT NULL,
		token VARCHAR(255) NOT NULL,
		refresh_token VARCHAR(255) NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
	if dbErr != nil {
		log.Fatalln(dbErr.Error())
	}
	log.Println("Connected to the database successfully")
	database = &Postgres{
		Db: db,
	}
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
