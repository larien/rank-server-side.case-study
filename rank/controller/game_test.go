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

func TestFindGamesByCategory(t *testing.T) {
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

	controller.StoreGame(g1)
	controller.StoreGame(g2)
	controller.StoreGame(g3)

	t.Run("should return all games from Puzzle category", func(t *testing.T) {
		category := "puzzle"
		games, err := controller.FindGamesByCategory(category)
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
		games, err := controller.FindGamesByCategory(category)
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
		games, err := controller.FindGamesByCategory(category)
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
		games, err := controller.FindGamesByCategory(category)
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
		games, err := controller.FindGamesByCategory(category)
		assert.Equal(t, 0, len(games))
		assert.Nil(t, err)
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

		game, err := controller.FindGameByID(id)
		assert.Equal(t, true, util.IsValidID(game.ID.String()))
		assert.Equal(t, game.Name, name)
		assert.Nil(t, err)
	})
}

func TestDeleteByIDGame(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	t.Run("should delete Game from inserted ID", func(t *testing.T) {
		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		id, _ := controller.StoreGame(g1) // TODO

		err := controller.DeleteGameByID(id)
		assert.Nil(t, err)

		game, errGetByID := controller.FindGameByID(id)
		assert.Nil(t, game)
		assert.Nil(t, errGetByID)
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

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

	t.Run("should update Game title", func(t *testing.T) {

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		id, _ := controller.StoreGame(g1) // TODO

		game, errGetByID := controller.FindGameByID(id)
		assert.Nil(t, errGetByID)
		assert.Equal(t, name, game.Name)

		differentName := "Game Name"

		game.Name = differentName
		err := controller.UpdateGame(game)
		assert.Nil(t, err)

		updatedGame, errGetByID2 := controller.FindGameByID(id)
		assert.Nil(t, errGetByID2)
		assert.Equal(t, differentName, updatedGame.Name)
	})
}

func TestFindAllCategories(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newGameController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.CATEGORY_COLLECTION).RemoveAll(nil)

	t.Run("should return all categories", func(t *testing.T) {
		if err := repo.InsertCategories(); err != nil {
			log.Fatal(err.Error())
		}

		categories, err := controller.FindAllCategories()
		assert.Nil(t, err)
		assert.True(t, len(categories) > 0)
	})

	t.Run("shouldnt return any category because they werent inserted", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.CATEGORY_COLLECTION).RemoveAll(nil)

		categories, err := controller.FindAllCategories()
		assert.Nil(t, err)
		assert.True(t, len(categories) == 0)
	})

}
