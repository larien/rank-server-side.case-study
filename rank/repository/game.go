package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Game defines the methods must be implemented by injected layer.
type Game interface {
	// DeleteGameByID(util.Identifier) error
	FindAllGames() ([]*entity.Game, error)
	// GetGameByID(util.Identifier) (*entity.Game, error)
	StoreGame(*entity.Game) (util.Identifier, error)
	// UpdateGame(*entity.Game) error
}

// FindAllGames returns all Game from the database sorted by ID.
func (m *MongoDB) FindAllGames() ([]*entity.Game, error) {
	var games []*entity.Game

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)
	if err := collection.Find(nil).Sort("id").All(&games); err != nil {
		return nil, err
	}

	return games, nil
}

// StoreGame inserts a new Game in the database.
func (m *MongoDB) StoreGame(game *entity.Game) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.GAME_COLLECTION)

	game.ID = util.NewID()

	coll.Insert(game)

	return game.ID, nil
}
