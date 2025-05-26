package db

import (
	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) AddProductIntoCart(cart models.Cart, product models.Prouduct, user models.User) error {
	stmt, err := db.Db.Prepare(`CREATE TABLE IF NOT EXISTS cart(
	iD        SERIAL  PRIMARY KEY,
	user_id    int   NOT NULL,
	product_id int   NOT NULL,
	quantity  int   NOT NULL,
	price     NUMERIC NOT NULL,
	total     NUMERIC,
	checkout  bool  DEFAULT false,
	createdAt date DEFAULT current_date,
	updatedAt date DEFAULT current_date,
	deletedAt date DEFAULT current_date,
	FOREGIN KEY(user_id) REFERENCES users(id)
	FOREGIN KEY(product_id) REFERENCES products(id)
	)`)
	if err != nil {
		return err
	}

	stmt.Exec()
	defer stmt.Close()

	insertStmt, err := db.Db.Prepare(`INSERT INTO cart (user_id,product_id,quantity,price,total,checkout,createdAt,updatedAt,deletedAt) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(user.ID, product.ID, cart.Quantity, cart.Price, cart.Total, cart.Checkout, cart.CreatedAt, cart.UpdatedAt, cart.DeletedAt)
	return err

}

func (db *Postgres) RemoveProductFromCart(productID int, userID int) error {
	stmt, err := db.Db.Prepare(`DELETE FROM cart WHERE product_id=$1 AND user_id=$2`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(productID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Postgres) GetItemFromCart(userID int) ([]models.Cart, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM cart WHERE user_id=$1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var carts []models.Cart
	res, err := stmt.Query(userID)
	if err != nil {
		return []models.Cart{}, err
	}
	for res.Next() {
		var cart models.Cart
		if err := res.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.Quantity, &cart.Price, &cart.Total, &cart.Checkout, &cart.CreatedAt, &cart.UpdatedAt, &cart.DeletedAt); err != nil {
			return []models.Cart{}, err
		}
		carts = append(carts, cart)
	}
	if err := res.Err(); err != nil {
		return []models.Cart{}, err
	}
	return carts, nil

}

func (db *Postgres) SearchProducts(query string) ([]models.Prouduct, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM products WHERE name ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%'`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var products []models.Prouduct
	res, err := stmt.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var product models.Prouduct
		if err := res.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.Stock, &product.Category, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
