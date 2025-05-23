package userdb

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/types"
)

func (db *Postgres) GetUserByEmail(email string) (types.User, error) {
	stmt, err := db.Db.Prepare("SELECT * FROM users WHERE email=?")
	if err != nil {
		return types.User{}, err
	}
	defer stmt.Close()
	var user types.User

	err = db.Db.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fmt.Errorf("no user found with email %s", email)
		}
		return types.User{}, err
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
