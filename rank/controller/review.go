package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review contains the injected Review interface from Repository layer.
type Review struct {
	Repository repository.Review
}

// ReviewController contains methods that must be implemented by the injected layer.
type ReviewController interface {
	FindAll() ([]*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
	GetByID(util.Identifier) (*entity.Review, error)
}

// newReviewController creates a new Review Controller.
func newReviewController(m *repository.MongoConn) *Review {
	return &Review{
		Repository: m,
	}
}

// FindAll returns all reviews from database.
func (r *Review) FindAll() ([]*entity.Review, error) {
	return r.Repository.FindAll()
}

// Store inserts a new Review in the database.
func (r *Review) Store(review *entity.Review) (util.Identifier, error) {
	return r.Repository.Store(review)
}

// GetByID inserts a new Review in the database.
func (r *Review) GetByID(id util.Identifier) (*entity.Review, error) {
	return r.Repository.GetByID(id)
}

// DeleteByID deletes a Review from the database by its ID.
func (r *Review) DeleteByID(id util.Identifier) error {
	return r.Repository.DeleteByID(id)
}
