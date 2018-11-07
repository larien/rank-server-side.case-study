package repository

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

func TestFindAll(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := NewMongoConnection(pool, config.MONGODB_DATABASE)

	t.Run("should have returned all reviews", func(t *testing.T) {
		// TODO - change to RemoveAll function from Repository layer
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		m.Store(r1)
		reviews, err := m.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
	})

	t.Run("should have returned error", func(t *testing.T) {
		m = NewMongoConnection(pool, "otherdatabase")
		reviews, err := m.FindAll()
		assert.NotNil(t, err)
		assert.Nil(t, reviews)
	})
}

func TestStore(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := NewMongoConnection(pool, config.MONGODB_DATABASE)

	t.Run("should have inserted a new review", func(t *testing.T) {
		// TODO - change to RemoveAll function from Repository layer
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		id := m.Store(r1)

		reviews, errFindAll := m.FindAll()

		assert.Equal(t, 1, len(reviews))
		assert.Nil(t, err)
		assert.Nil(t, errFindAll)
		assert.NotNil(t, id)
		assert.Equal(t, true, util.IsValidID(id.String()))
	})
}
