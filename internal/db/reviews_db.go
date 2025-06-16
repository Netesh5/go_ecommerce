package db

import (
	"fmt"

	"github.com/netesh5/go_ecommerce/internal/models"
)

func (db *Postgres) AddReview(review models.Review) error {
	stmt, err := db.Db.Prepare(`INSERT INTO reviews 
	(user_id,product_id,rating,comment,created_at,updated_at) 
	VALUES($1,$2,$3,$4,$5,$6)
	`)
	if err != nil {
		return fmt.Errorf("couldn't insert data into tabble")
	}

	_, err = stmt.Exec(&review.UserID, &review.ProductID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (db *Postgres) GetProductReviews(productId int) ([]models.Review, error) {
	stmt, err := db.Db.Prepare(`SELECT * FROM reviews WHERE product_id=$1`)
	if err != nil {
		return []models.Review{}, err
	}
	reviews := []models.Review{}
	defer stmt.Close()
	res, err := stmt.Query(productId)
	if err != nil {
		return []models.Review{}, err
	}

	for res.Next() {
		var review models.Review
		if err := res.Scan(&review.Id, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt); err != nil {
			return []models.Review{}, err
		}

		reviews = append(reviews, review)
	}
	return reviews, nil
}
