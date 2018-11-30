package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	"gopkg.in/mgo.v2/bson"
)

// Game defines the methods must be implemented by injected layer.
type Game interface {
	DeleteGameByID(util.Identifier) error
	FindAllGames() ([]*entity.Game, error)
	FindGameByID(util.Identifier) (*entity.Game, error)
	FindGamesByCategory(category string) ([]*entity.Game, error)
	StoreGame(*entity.Game) (util.Identifier, error)
	UpdateGame(*entity.Game) error
	FindAllCategories() ([]string, error)
}

// DeleteGameByID deletes a Game by its ID.
func (m *MongoDB) DeleteGameByID(id util.Identifier) error {
	return m.pool.Session(nil).DB(m.db).C(config.GAME_COLLECTION).RemoveId(id)
}

// FindAllGames returns all Games from the database sorted by ID.
func (m *MongoDB) FindAllGames() ([]*entity.Game, error) {
	var games []*entity.Game

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)
	if err := collection.Find(nil).Sort("id").All(&games); err != nil {
		return nil, err
	}

	return games, nil
}

// FindGamesByCategory returns all Games from the database filtered by category.
func (m *MongoDB) FindGamesByCategory(category string) ([]*entity.Game, error) {
	var games []*entity.Game
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)
	if err := collection.Find(bson.M{"categories": category}).All(&games); err != nil {
		return nil, err
	}

	return games, nil
}

// FindGameByID finds a Game by its ID.
func (m *MongoDB) FindGameByID(id util.Identifier) (*entity.Game, error) {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)

	var game *entity.Game

	collection.FindId(id).One(&game)

	return game, nil
}

// StoreGame inserts a new Game in the database.
func (m *MongoDB) StoreGame(game *entity.Game) (util.Identifier, error) {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)

	game.ID = util.NewID()

	collection.Insert(game)

	return game.ID, nil
}

// UpdateGame updates an existing Game in the database.
func (m *MongoDB) UpdateGame(game *entity.Game) error {
	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.GAME_COLLECTION)

	_, err := collection.UpsertId(game.ID, game) // TODO - avoid null Games
	return err
}

// InsertCategories insert category names in the Category collection from database.
func (m *MongoDB) InsertCategories() error {

	var c entity.Categories

	c.Names = []string{
		"action",
		"adventure",
		"fighting",
		"first-person shooting",
		"puzzle",
		"racing",
		"real-time strategy",
		"sports",
		"role-playing",
		"third-person shooting",
	}

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.CATEGORY_COLLECTION)

	collection.RemoveAll(nil)

	return collection.Insert(&c)
}

// FindAllCategories returns all categories from database.
func (m *MongoDB) FindAllCategories() ([]string, error) {
	var category []entity.Categories

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.CATEGORY_COLLECTION)
	if err := collection.Find(nil).All(&category); err != nil {
		return nil, err
	}

	return category[0].Names, nil
}
