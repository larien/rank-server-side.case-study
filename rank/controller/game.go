package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Game contains the injected Game interface from Repository layer.
type Game struct {
	Repository repository.Game
}

// GameController contains methods that must be implemented by the injected layer.
type GameController interface {
	DeleteGameByID(util.Identifier) error
	FindAllGames() ([]*entity.Game, error)
	FindGamesByCategory(category string) ([]*entity.Game, error)
	FindGameByID(util.Identifier) (*entity.Game, error)
	StoreGame(*entity.Game) (util.Identifier, error)
	UpdateGame(*entity.Game) error
}

// newGameController creates a new Game Controller.
func newGameController(m *repository.MongoDB) *Game {
	return &Game{
		Repository: m,
	}
}

// DeleteGameByID requests the Repository layer for a Game to be deleted from the database by its ID.
func (g *Game) DeleteGameByID(id util.Identifier) error {
	return g.Repository.DeleteGameByID(id)
}

// FindAllGames requests the Repository layer to return all Games from database.
func (g *Game) FindAllGames() ([]*entity.Game, error) {
	return g.Repository.FindAllGames()
}

// FindGamesByCategory requests the Repository layer to return all Games from database
// filtering by category.
func (g *Game) FindGamesByCategory(category string) ([]*entity.Game, error) {
	return g.Repository.FindGamesByCategory(category)
}

// FindGameByID requests the Repository layer for a certain Game by its ID.
func (g *Game) FindGameByID(id util.Identifier) (*entity.Game, error) {
	return g.Repository.FindGameByID(id)
}

// StoreGame requests the Repository layer for the insertion of a new Game in the database.
func (g *Game) StoreGame(game *entity.Game) (util.Identifier, error) {
	return g.Repository.StoreGame(game)
}

// UpdateGame requests the Repository layer for a Game to be updated in the database.
func (g *Game) UpdateGame(game *entity.Game) error {
	return g.Repository.UpdateGame(game)
}
