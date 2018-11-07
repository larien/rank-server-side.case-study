package repository

import (
	"github.com/juju/mgosession"
)

// MongoConn mongodb repo
type MongoConn struct {
	pool *mgosession.Pool
	db   string
}

// NewMongoConnection create new repository
func NewMongoConnection(p *mgosession.Pool, db string) *MongoConn {
	return &MongoConn{
		pool: p,
		db:   db,
	}
}
