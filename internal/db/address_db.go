package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) DeleteAddress(id int, userId int) error {
	stmt, err := db.Db.Prepare(`DELETE FROM addresses WHERE id=$1 AND user_id=$2`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, userId)
	if err != nil {
		return err
	}
	return nil

}

func (db *Postgres) GetUserAddress(id int, userId int) (models.Address, error) {

	var address models.Address
	stmt, dbError := db.Db.Prepare(`SELECT * FROM addresses where id=$1 AND user_id=$2`)
	if dbError != nil {
		return models.Address{}, dbError
	}
	qErr := stmt.QueryRow(id, userId).Scan(&address.Id, &address.Address, &address.City, &address.State, &address.Country, &address.ZipCode, &address.UserId, &address.CreatedAt, &address.UpdatedAt)
	if qErr != nil {
		return models.Address{}, qErr
	}
	return address, nil

}

func (db *Postgres) GetUserAddresses(userId int) ([]models.Address, error) {
	var addresses []models.Address

	stmt, err := db.Db.Prepare(`SELECT * from addresses WHERE user_id=$1`)
	if err != nil {
		return []models.Address{}, nil
	}

	res, err := stmt.Query(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return []models.Address{}, nil
		}
		return []models.Address{}, err
	}

	for res.Next() {
		var address models.Address
		err := res.Scan(&address.Id, &address.Address, &address.City, &address.State, &address.Country, &address.ZipCode, &address.UserId, &address.CreatedAt, &address.UpdatedAt)
		if err != nil {
			return []models.Address{}, err
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (db *Postgres) AddUserAddress(address models.Address) error {
	stmt, err := db.Db.Prepare(`UPDATE addresses SET address=$1,city=$2,state=$3,country=$4,zipcode=$5,user_id=$6,updated_at=7`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(address.Address, address.City, address.State, address.Country, address.ZipCode, address.UserId, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("couldn't add user address")
	}
	return nil
}
