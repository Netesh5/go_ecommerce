package db

import "github.com/netesh5/go_ecommerce/internal/models"

func (db *Postgres) AddProductToWishList(wishlist models.Wishlists) error {
	stmt, err := db.Db.Prepare(`INSERT INTO wishlists (user_id,product_id,created_at,updated_at)VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&wishlist.UserId, &wishlist.ProductId, &wishlist.CreatedAt, &wishlist.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (db *Postgres) GetUserWishlistProducts(userId int) ([]models.Wishlists, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM wishlists WHERE user_id=$1`)
	if err != nil {
		return []models.Wishlists{}, err
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return []models.Wishlists{}, err
	}
	defer stmt.Close()
	wishlists := []models.Wishlists{}
	for res.Next() {
		var wishlist models.Wishlists

		if err := res.Scan(&wishlist.Id, &wishlist.UserId, &wishlist.ProductId, &wishlist.CreatedAt, &wishlist.UpdatedAt); err != nil {
			return nil, err
		}
		wishlists = append(wishlists, wishlist)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return wishlists, nil

}
