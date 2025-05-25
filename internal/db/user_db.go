package db

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetUserByEmail(email string) (models.User, error) {
	stmt, err := db.Db.Prepare("SELECT id, email FROM users WHERE email= $1")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with email %s", email)
		}
		return models.User{}, err
	}
	return user, nil

}

//	func GetUserByID(id int) (types.User, error) {
//		// Implementation to get user by ID
//	}
func (db *Postgres) CreateUser(user models.User) (models.User, error) {
	res, err := db.Db.Exec(`INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	user.ID = int(lastInsertID)
	return user, nil
}

// func UpdateUser(user types.User) (types.User, error) {
// 	// Implementation to update user details
// }
// func DeleteUser(id int) error {
// 	// Implementation to delete a user by ID
// }

func (db *Postgres) GetUserByID(id int) (models.User, error) {
	stmt, err := db.Db.Prepare("SELECT id, name, email FROM users WHERE id = $1")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with ID %d", id)
		}
		return models.User{}, err
	}
	return user, nil
}
