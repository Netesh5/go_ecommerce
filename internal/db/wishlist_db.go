package db

func (db *Postgres) AddProductToWishList(productId int) error {
	stmt, err := db.Db.Prepare(`INSERT INTO TABLE (user_id,product_id,created_at,updated_at)VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
}
