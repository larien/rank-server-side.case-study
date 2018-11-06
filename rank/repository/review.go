package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

type ReviewRepository interface {
	FindAll() ([]*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
}

func (m *MongoConn) FindAll() ([]*entity.Review, error) {
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

func (r *MongoConn) Store(v *entity.Review) (util.Identifier, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("review")
	err := coll.Insert(v)
	if err != nil {
		return "", err
	}
	return v.ID, nil
}
