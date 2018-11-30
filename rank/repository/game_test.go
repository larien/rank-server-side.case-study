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

// TODO - create Store Game tests

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
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		m = New(pool, "otherdatabase")
		games, err := m.FindAllGames()
		assert.NotNil(t, err)
		assert.Nil(t, games)
	})
}

func TestFindGamesByCategory(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

	g1 := &entity.Game{
		Name:       "First game",
		Categories: []string{"action", "puzzle"},
	}

	g2 := &entity.Game{
		Name:       "Second game",
		Categories: []string{"adventure", "puzzle"},
	}

	g3 := &entity.Game{
		Name:       "Third game",
		Categories: []string{"adventure", "third-person shooter"},
	}

	m.StoreGame(g1)
	m.StoreGame(g2)
	m.StoreGame(g3)

	t.Run("should return all games from Puzzle category", func(t *testing.T) {
		category := "puzzle"
		games, err := m.FindGamesByCategory(category)
		assert.Equal(t, 2, len(games))
		assert.Nil(t, err)

		foundFirstGame := false
		for _, game := range games {
			if game.Name == g1.Name {
				foundFirstGame = true
			}
		}
		assert.True(t, foundFirstGame)

		foundSecondGame := false
		for _, game := range games {
			if game.Name == g2.Name {
				foundSecondGame = true
			}
		}
		assert.True(t, foundSecondGame)
	})

	t.Run("should return all games from Adventure category", func(t *testing.T) {
		category := "adventure"
		games, err := m.FindGamesByCategory(category)
		assert.Equal(t, 2, len(games))
		assert.Nil(t, err)

		foundSecondGame := false
		for _, game := range games {
			if game.Name == g2.Name {
				foundSecondGame = true
			}
		}
		assert.True(t, foundSecondGame)

		foundThirdGame := false
		for _, game := range games {
			if game.Name == g3.Name {
				foundThirdGame = true
			}
		}
		assert.True(t, foundThirdGame)
	})

	t.Run("should return all games from Third-person Shooter category", func(t *testing.T) {
		category := "third-person shooter"
		games, err := m.FindGamesByCategory(category)
		assert.Equal(t, 1, len(games))
		assert.Nil(t, err)

		foundThirdGame := false
		for _, game := range games {
			if game.Name == g3.Name {
				foundThirdGame = true
			}
		}
		assert.True(t, foundThirdGame)
	})

	t.Run("should return all games from Action category", func(t *testing.T) {
		category := "action"
		games, err := m.FindGamesByCategory(category)
		assert.Equal(t, 1, len(games))
		assert.Nil(t, err)

		foundFirstGame := false
		for _, game := range games {
			if game.Name == g1.Name {
				foundFirstGame = true
			}
		}
		assert.True(t, foundFirstGame)
	})

	t.Run("shouldnt return any games from inexistent category", func(t *testing.T) {
		category := "none"
		games, err := m.FindGamesByCategory(category)
		assert.Equal(t, 0, len(games))
		assert.Nil(t, err)
	})

	t.Run("should have returned error", func(t *testing.T) {

		m = New(pool, "otherdatabase")
		games, err := m.FindGamesByCategory("lalala")
		assert.NotNil(t, err)
		assert.Nil(t, games)
	})
}

func TestFindGameByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should find certain Game by stored ID", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, _ := m.StoreGame(g1)

		game, err := m.FindGameByID(id)
		assert.Equal(t, name, game.Name)
		assert.Nil(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, game.ID)
		assert.Equal(t, true, util.IsValidID(id.String()))
		assert.Equal(t, true, util.IsValidID(game.ID.String()))
	})
}

func TestDeleteGameByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should delete certain Game by stored ID", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, _ := m.StoreGame(g1)

		game, errGetByID := m.FindGameByID(id)
		assert.Equal(t, id, game.ID)
		assert.Nil(t, errGetByID)

		err := m.DeleteGameByID(id)
		assert.Nil(t, err)

		game, errGetByID2 := m.FindGameByID(id)
		assert.Nil(t, game)
		assert.Nil(t, errGetByID2)
	})
}

func TestUpdateGame(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have updated a new game", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, err := m.StoreGame(g1)
		assert.Nil(t, err)

		game, errGetByID := m.FindGameByID(id)
		assert.Nil(t, errGetByID)
		assert.Equal(t, name, game.Name)

		differentName := "Different name"

		game.Name = differentName
		errUpdate := m.UpdateGame(game)
		assert.Nil(t, errUpdate)

		updatedGame, errGetByID2 := m.FindGameByID(id)

		assert.Nil(t, errGetByID2)
		assert.Equal(t, differentName, updatedGame.Name)
	})
}
