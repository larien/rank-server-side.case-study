package controller

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
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

	repo := repository.NewMongoConnection(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	// TODO - change to RemoveAll function from Repository layer
	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	r1 := &entity.Review{
		Title: "Title 1",
	}

	controller.Store(r1)

	t.Run("should return inserted review with 'Title 1' as title", func(t *testing.T) {
		reviews, err := controller.FindAll()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
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

	repo := repository.NewMongoConnection(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	// TODO - change to RemoveAll function from Repository layer
	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	r1 := &entity.Review{
		Title: "Title 1",
	}

	t.Run("should return inserted ID", func(t *testing.T) {
		id := controller.Store(r1)
		assert.Equal(t, true, util.IsValidID(id.String()))
	})

	t.Run("should have inserted new review", func(t *testing.T) {
		reviews, errFindAll := controller.FindAll()
		assert.Nil(t, errFindAll)
		assert.Equal(t, 1, len(reviews))
	})
}
