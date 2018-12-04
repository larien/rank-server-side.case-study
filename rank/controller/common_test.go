package controller

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/middlewares/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	mgo "gopkg.in/mgo.v2"
)

func TestNew(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	t.Run("should have created Review controller", func(t *testing.T) {

		controllers := New(repo)
		assert.NotNil(t, controllers.Review)
	})
}
