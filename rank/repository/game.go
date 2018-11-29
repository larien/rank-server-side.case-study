package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Game defines the methods must be implemented by injected layer.
type Game interface {
	DeleteGameByID(util.Identifier) error
	FindAllGames() ([]*entity.Game, error)
	GetGameByID(util.Identifier) (*entity.Game, error)
	StoreGame(*entity.Game) (util.Identifier, error)
	UpdateGame(*entity.Game) error
}

// DeleteGameByID deletes a Game by its ID.
func (m *MongoDB) DeleteGameByID(id util.Identifier) error {
	return m.pool.Session(nil).DB(m.db).C(config.GAME_COLLECTION).RemoveId(id)
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

// GetGameByID finds a Game by its ID.
func (m *MongoDB) GetGameByID(id util.Identifier) (*entity.Game, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.GAME_COLLECTION)

	var game *entity.Game

	coll.FindId(id).One(&game)

	return game, nil
}

// StoreGame inserts a new Game in the database.
func (m *MongoDB) StoreGame(game *entity.Game) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.GAME_COLLECTION)

	game.ID = util.NewID()

	coll.Insert(game)

	return game.ID, nil
}

// UpdateGame updates an existing Game in the database.
func (m *MongoDB) UpdateGame(game *entity.Game) error {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.GAME_COLLECTION)

	_, err := coll.UpsertId(game.ID, game) // TODO - avoid null Games
	return err
}
