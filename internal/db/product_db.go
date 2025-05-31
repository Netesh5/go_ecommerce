package db

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetProductByID(id int) (models.Product, error) {
	stmt, err := db.Db.Prepare("SELECT * from product WHERE id=$1")
	if err != nil {
		return models.Product{}, err
	}

	defer stmt.Close()
	var product models.Product
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
			return models.Product{}, fmt.Errorf("no product found with")
		}
		return models.Product{}, err
	}
	return product, nil
}

func (db *Postgres) GetAllProducts() ([]models.Product, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM products`)
	if err != nil {
		return []models.Product{}, err
	}

	defer stmt.Close()
	var products []models.Product
	res, err := stmt.Query()
	if err != nil {
		return []models.Product{}, err
	}
	for res.Next() {
		var product models.Product
		if err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}
	if err := res.Err(); err != nil {
		return []models.Product{}, err
	}
	return products, nil

}
