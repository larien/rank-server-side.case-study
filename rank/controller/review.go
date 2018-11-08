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
	DeleteByID(util.Identifier) error
	FindAll() ([]*entity.Review, error)
	GetByID(util.Identifier) (*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
	Update(*entity.Review) error
}

// newReviewController creates a new Review Controller.
func newReviewController(m *repository.MongoDB) *Review {
	return &Review{
		Repository: m,
	}
}

// DeleteByID requests the Repository layer for a Review to be deleted from the database by its ID.
func (r *Review) DeleteByID(id util.Identifier) error {
	return r.Repository.DeleteByID(id)
}

// FindAll requests the Repository layer to return all Reviews from database.
func (r *Review) FindAll() ([]*entity.Review, error) {
	return r.Repository.FindAll()
}

// GetByID requests the Repository layer for a certain Review by its ID.
func (r *Review) GetByID(id util.Identifier) (*entity.Review, error) {
	return r.Repository.GetByID(id)
}

// Store requests the Repository layer for the insertion of a new Review in the database.
func (r *Review) Store(review *entity.Review) (util.Identifier, error) {
	return r.Repository.Store(review)
}

// Update requests the Repository layer for a Review to be updated in the database.
func (r *Review) Update(review *entity.Review) error {
	return r.Repository.Update(review)
}
