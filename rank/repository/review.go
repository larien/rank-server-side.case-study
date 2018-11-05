package repository

import (
	"github.com/juju/mgosession"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

//NewReviewRepository create new repository
func NewReviewRepository(p *mgosession.Pool, db string) ReviewRepository {
	return &mongo{
		pool: p,
		db:   db,
	}
}

func (m *mongo) FindAll() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C("reviews")
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
