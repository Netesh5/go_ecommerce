package db

import (
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) AddReview(review models.ReviewRequest) error {
	stmt, err := db.Db.Prepare(`INSERT INTO reviews 
	(id,user_id,product_id,rating,comment,created_at,updated_at) 
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`)
	if err != nil {
		return fmt.Errorf("couldn't insert data into tabble")
	}
}
