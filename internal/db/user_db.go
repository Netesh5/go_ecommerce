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
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.Phone, &user.Token, &user.RefreshToken, &user.Address, &user.Cart, &user.Orders, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with email %s", email)
		}
		return models.User{}, err
	}
	return user, nil

}

func (db *Postgres) CreateUser(user models.User) (models.User, error) {
	// res, err := db.Db.Exec(`INSERT INTO users (name, email, password,phone,token,refresh_token, created_at, updated_at) VALUES ($1, $2, $3, $4, $5,$6,$7,$8)`, user.Name, user.Email, user.Password, user.Phone, user.Token, user.RefreshToken, user.CreatedAt, user.UpdatedAt)
	// if err != nil {
	// 	return models.User{}, err
	// }
	// lastInsertID, err := res.LastInsertId()
	// if err != nil {
	// 	return models.User{}, err
	// }
	// user.ID = int(lastInsertID)
	// return user, nil
	err := db.Db.QueryRow(`
    INSERT INTO users (name, email, password, phone, token, refresh_token, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
`, user.Name, user.Email, user.Password, user.Phone, user.Token, user.RefreshToken, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

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

func (db *Postgres) UpdateUser(user models.User) (models.User, error) {
	stmt, err := db.Db.Prepare(`UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4,token=$5,refresh_token=$6,phone=$7,address=$8,cart=$9,orders=$10 WHERE id = $11`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.UpdatedAt, user.Token, user.RefreshToken, user.Phone, user.Address, user.Cart, user.Orders, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
