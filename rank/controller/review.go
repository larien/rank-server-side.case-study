package controller

import (
	"fmt"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review data
type Review struct {
	Repository repository.ReviewRepository
}

type ReviewController interface {
	FindAll() ([]*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
}

// NewReviewController creates a new Review Controller.
func NewReviewController(c *repository.MongoConn) *Review {
	return &Review{
		Repository: c,
	}
}

// FindAll returns all reviews from database.
func (r *Review) FindAll() ([]*entity.Review, error) {
	return r.Repository.FindAll()
}

func (r *Review) Store(review *entity.Review) (util.Identifier, error) {
	fmt.Printf("Review recebida 2: %+v", review)
	return r.Repository.Store(review)
}
