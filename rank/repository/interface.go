package repository

import (
	"github.com/juju/mgosession"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

// mongo mongodb repo
type mongo struct {
	pool *mgosession.Pool
	db   string
}

// ReviewRepository interface defines the function Review must implement.
type ReviewRepository interface {
	FindAll() ([]*entity.Review, error)
}
