package db

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetUserByEmail(email string) (models.User, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM users WHERE email = $1`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	var userAddess []models.Address
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Token, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// For signup checks, this means the email is not registered yet
			return models.User{}, nil
		}
		return models.User{}, err
	}

	addSmt, err := db.Db.Prepare(`SELECT * FROM addresses where id =$1`)
	if err != nil {
		return models.User{}, err
	}

	defer addSmt.Close()

	res, err := addSmt.Query(user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			// For signup checks, this means the email is not registered yet
			return models.User{}, nil
		}
		return models.User{}, err
	}

	for res.Next() {
		var addr models.Address

		if err := res.Scan(&addr.Id, &addr.UserId, &addr.Address, &addr.City, &addr.Country, &addr.State, &addr.ZipCode, &addr.CreatedAt, &addr.UpdatedAt); err != nil {
			return models.User{}, err
		}

		userAddess = append(userAddess, addr)
	}

	user.Address = userAddess

	return user, nil
}

func (db *Postgres) CreateUser(user models.User) (models.User, error) {
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
	stmt, err := db.Db.Prepare(`UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4,token=$5,refresh_token=$6,phone=$7 WHERE id = $8`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.UpdatedAt, user.Token, user.RefreshToken, user.Phone, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
