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
		is_verified BOOLEAN DEFAULT FALSE,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		); 

		
		CREATE TABLE IF NOT EXISTS addresses (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		street TEXT NOT NULL,
		city TEXT NOT NULL,
		state TEXT,
		country TEXT NOT NULL,
		postal_code TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

		CREATE TABLE IF NOT EXISTS products(
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		price NUMERIC NOT NULL,
		image VARCHAR(255) NOT NULL,
		stock NUMERIC NOT NULL,
		category VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);


		CREATE TABLE IF NOT EXISTS cart(
		id      SERIAL PRIMARY KEY,
		user_id    int   NOT NULL,
		product_id int   NOT NULL,
		quantity  int   NOT NULL,
		price     NUMERIC NOT NULL,
		total     NUMERIC NOT NULL, 
		checkout  bool  DEFAULT false,
		createdAt date DEFAULT current_date,
		updatedAt date DEFAULT current_date,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(product_id) REFERENCES products(id)
	);

	CREATE TABLE IF NOT EXISTS reviews(
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		product_id INTEGER REFERENCES products(id),
		rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
		comment TEXT,
		created_at TIMESTAMP DEFAULT current_date,
		updated_at TIMESTAMP DEFAULT current_date
	);

	CREATE TABLE IF NOT EXISTS wishlists (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		product_id INTEGER REFERENCES products(id),
		created_at TIMESTAMP DEFAULT current_date,
		updated_at TIMESTAMP DEFAULT current_date,
		UNIQUE (user_id, product_id)
	);

		`)

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
