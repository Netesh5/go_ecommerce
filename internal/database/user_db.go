package userdb

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetUserByEmail(email string) (models.User, error) {
	stmt, err := db.Db.Prepare("SELECT id, email FROM users WHERE email=$1")
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

// func GetUserByID(id int) (types.User, error) {
// 	// Implementation to get user by ID
// }
// func CreateUser(user types.User) (types.User, error) {
// 	// Implementation to create a new user
// }
// func UpdateUser(user types.User) (types.User, error) {
// 	// Implementation to update user details
// }
// func DeleteUser(id int) error {
// 	// Implementation to delete a user by ID
// }
