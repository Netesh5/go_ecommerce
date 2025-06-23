package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetUserByEmail(email string) (models.User, error) {
	stmt, err := db.Db.Prepare(`SELECT id,name,email,password,phone,token,refresh_token,created_at,updated_at,is_verified FROM users WHERE email = $1`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	var userAddess []models.Address
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Token, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt, &user.Verfiy)
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

		if err := res.Scan(&addr.Id, &addr.UserId, &addr.Street, &addr.City, &addr.Country, &addr.State, &addr.ZipCode, &addr.CreatedAt, &addr.UpdatedAt); err != nil {
			return models.User{}, err
		}

		userAddess = append(userAddess, addr)
	}

	user.Address = userAddess

	return user, nil
}

func (db *Postgres) CreateUser(user models.User) (models.User, error) {
	err := db.Db.QueryRow(`
    INSERT INTO users (name, email, password, phone, token, refresh_token,is_verified, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9)
    RETURNING id
`, user.Name, user.Email, user.Password, user.Phone, user.Token, user.RefreshToken, user.Verfiy, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)

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
	stmt, err := db.Db.Prepare(`UPDATE users SET name = $1, email = $2, updated_at = $3,token=$4,refresh_token=$5,phone=$6,is_verified=$7 WHERE id = $8`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.UpdatedAt, user.Token, user.RefreshToken, user.Phone, user.Verfiy, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (db *Postgres) UpdateUserInfo(user models.UserUpdate) (models.UserUpdate, error) {
	stmt, err := db.Db.Prepare(`UPDATE users SET name = $1, email = $2, updated_at = $3,phone=$4 WHERE id = $5`)
	if err != nil {
		return models.UserUpdate{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, time.Now().UTC(), user.Phone, user.ID)
	if err != nil {
		return models.UserUpdate{}, err
	}

	return user, nil
}
func (db *Postgres) UpdateToken(user models.User) (models.User, error) {
	stmt, err := db.Db.Prepare(`UPDATE users SET token=$1,refresh_token=$2 WHERE id = $3`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Token, user.RefreshToken, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (db *Postgres) MarkUserVerified(userID int) error {
	_, err := db.Db.Exec(`UPDATE users SET is_verified = TRUE, updated_at = $1 WHERE id = $2`, time.Now().UTC(), userID)
	return err
}

func (db *Postgres) CheckEmail(email string) error {
	stmt, err := db.Db.Prepare(`SELECT email from users WHERE email=$1`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var user models.User

	err = stmt.QueryRow(email).Scan(&user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (db *Postgres) UpdatePassword(email string, password string) error {

	user, err := db.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	stmt, err := db.Db.Prepare(`UPDATE users SET password=$1,updated_at = $2 WHERE id=$3`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(password, time.Now().UTC(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}
