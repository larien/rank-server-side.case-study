package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
)

// Controllers contains the Controllers for each Entity.
type Controllers struct {
	Review repository.Review
	Game   repository.Game
}

// New creates new Controllers for each Entity.
func New(repo *repository.MongoDB) *Controllers {
	return &Controllers{
		Review: newReviewController(repo),
		Game:   newGameController(repo),
	}
}
