package db

import "github.com/netesh5/go_ecommerce/internal/models"

func (db *Postgres) AddProductToWishList(wishlist models.Wishlists) error {
	stmt, err := db.Db.Prepare(`INSERT INTO TABLE (user_id,product_id,created_at,updated_at)VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&wishlist.UserId, &wishlist.ProductId, &wishlist.CreatedAt, &wishlist.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
