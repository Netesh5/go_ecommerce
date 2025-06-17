package db

import (
	"database/sql"
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) GetProductByID(id int) (models.Product, error) {
	stmt, err := db.Db.Prepare("SELECT * from products WHERE id=$1")
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
		&product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("no product found with")
		}
		return models.Product{}, err
	}
	return product, nil
}

func (db *Postgres) GetAllProducts(limit int, offset int) ([]models.Product, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM products LIMIT $1 OFFSET $2`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	products := []models.Product{}
	res, err := stmt.Query(limit, offset)

	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var product models.Product
		if err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return products, nil

}

func (db *Postgres) AddProduct(product models.Product) error {
	stmt, err := db.Db.Prepare(`INSERT INTO products (name,description,price,image,stock,category,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`)

	if err != nil {
		// return fmt.Errorf("failed to create product")
		return err
	}

	_, err = stmt.Exec(&product.Name, &product.Description, &product.Price, &product.Image, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		// return fmt.Errorf("failed to create product")
		return err
	}

	return nil
}

func (db *Postgres) GetProductCount() (int, error) {
	var count int
	err := db.Db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
