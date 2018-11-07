package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

// Review defines the methods must be implemented by injected layer.
type Review interface {
	FindAll() ([]*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
}

// FindAll returns all Reviews from the database sorted by ID.
func (m *MongoConn) FindAll() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.REVIEW_COLLECTION)
	err := collection.Find(nil).Sort("id").All(&reviews)
	switch err {
	case nil:
		return reviews, nil
	case mgo.ErrNotFound:
		return nil, util.ErrNotFound
	default:
		return nil, err
	}
}

// Store inserts a new Review in the database.
func (m *MongoConn) Store(review *entity.Review) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	review.ID = util.NewID()
	err := coll.Insert(review)
	if err != nil {
		return "", err
	}
	return review.ID, nil
}
