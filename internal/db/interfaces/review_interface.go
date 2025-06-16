package interfaces

import "github.com/netesh5/go_ecommerce/internal/models"

type IReviews interface {
	AddReview(models.Review) error
	GetReview() ([]models.ReviewRequest, error)
}
