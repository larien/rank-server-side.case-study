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

func TestFindAllGames(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

	name := "Game 1"

	g1 := &entity.Game{
		Name: name,
	}

	controller.StoreGame(g1)

	t.Run("should return inserted game with 'Game 1' as title", func(t *testing.T) {
		games, err := controller.FindAllGames()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(games))
		assert.Equal(t, name, games[0].Name)
	})
}

func TestStoreGame(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

	g1 := &entity.Game{
		Name: "Game Name",
	}

	t.Run("should return inserted ID", func(t *testing.T) {
		id, _ := controller.StoreGame(g1) // TODO
		assert.Equal(t, true, util.IsValidID(id.String()))
	})

	t.Run("should have inserted new game", func(t *testing.T) {
		games, errFindAll := controller.FindAllGames()
		assert.Nil(t, errFindAll)
		assert.Equal(t, 1, len(games))
	})
}

func TestGetByIDGame(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	t.Run("should return Game from inserted ID", func(t *testing.T) {
		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		id, _ := controller.StoreGame(g1) // TODO
		assert.Equal(t, true, util.IsValidID(id.String()))

		game, err := controller.GetGameByID(id)
		assert.Equal(t, true, util.IsValidID(game.ID.String()))
		assert.Equal(t, game.Name, name)
		assert.Nil(t, err)
	})
}
