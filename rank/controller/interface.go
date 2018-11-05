package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

// Service data
type Service struct {
	repo ReviewRepository
}

// ReviewRepository interface defines the function Review controller must implement.
type ReviewRepository interface {
	FindAll() ([]*entity.Review, error)
}
