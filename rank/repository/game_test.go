package repository

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	mgo "gopkg.in/mgo.v2"
)

func TestFindAllGames(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have returned all games", func(t *testing.T) {

		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game 1"

		g1 := &entity.Game{
			Name: name,
		}

		m.StoreGame(g1)
		games, err := m.FindAllGames()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(games))
		assert.Equal(t, name, games[0].Name)
	})

	t.Run("should have returned error", func(t *testing.T) {
		m = New(pool, "otherdatabase")
		games, err := m.FindAllGames()
		assert.NotNil(t, err)
		assert.Nil(t, games)
	})
}
