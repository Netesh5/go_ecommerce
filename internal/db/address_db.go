package db

import (
	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) DeleteAddress(id int, userId int) error {
	stmt, err := db.Db.Prepare(`DELETE FROM address WHERE id=$1 AND user_id=$2`)
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
	stmt, dbError := db.Db.Prepare(`SELECT * FROM address where id=$1 AND user_id=$2`)
	if dbError != nil {
		return models.Address{}, dbError
	}
	qErr := stmt.QueryRow(id, userId).Scan(&address.Id, &address.Address, &address.City, &address.State, &address.Country, &address.ZipCode, &address.UserId, &address.CreatedAt, &address.UpdatedAt)
	if qErr != nil {
		return models.Address{}, qErr
	}
	return address, nil

}
