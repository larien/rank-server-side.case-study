package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review defines the methods must be implemented by injected layer.
type Review interface {
	FindAll() ([]*entity.Review, error)
	Store(*entity.Review) util.Identifier
}

// FindAll returns all Reviews from the database sorted by ID.
func (m *MongoConn) FindAll() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.REVIEW_COLLECTION)
	if err := collection.Find(nil).Sort("id").All(&reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

// Store inserts a new Review in the database.
func (m *MongoConn) Store(review *entity.Review) util.Identifier {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	review.ID = util.NewID()

	coll.Insert(review)

	return review.ID
}
