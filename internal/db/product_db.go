package db

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetProductByID(id int) (models.Prouduct, error) {
	stmt, err := db.Db.Prepare("SELECT * from product WHERE id=$1")
	if err != nil {
		return models.Prouduct{}, err
	}

	defer stmt.Close()
	var product models.Prouduct
	err = stmt.QueryRow(id).Scan(&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.Stock,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Prouduct{}, fmt.Errorf("no product found with")
		}
		return models.Prouduct{}, err
	}
	return product, nil
}
