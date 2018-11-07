package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
)

// Controllers contains the controllers for each entity.
type Controllers struct {
	Review repository.ReviewRepository
}

// New creates new controllers.
func New(repo *repository.MongoConn) *Controllers {
	return &Controllers{
		Review: newReviewController(repo),
	}
}
