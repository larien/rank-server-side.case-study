package repository

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	mgo "gopkg.in/mgo.v2"
)

func TestDB(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	t.Run("should have defined MongoDB connection", func(t *testing.T) {
		m := NewMongoConnection(pool, config.MONGODB_DATABASE)

		assert.NotNil(t, m)
		assert.Equal(t, "rank", m.db)
	})
}
