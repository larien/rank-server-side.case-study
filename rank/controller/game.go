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
	// DeleteGameByID(util.Identifier) error
	FindAllGames() ([]*entity.Game, error)
	// GetGameByID(util.Identifier) (*entity.Game, error)
	StoreGame(*entity.Game) (util.Identifier, error)
	// UpdateGame(*entity.Game) error
}

// newGameController creates a new Game Controller.
func newGameController(m *repository.MongoDB) *Game {
	return &Game{
		Repository: m,
	}
}

// FindAllGames requests the Repository layer to return all Games from database.
func (r *Game) FindAllGames() ([]*entity.Game, error) {
	return r.Repository.FindAllGames()
}

// StoreGame requests the Repository layer for the insertion of a new Game in the database.
func (r *Game) StoreGame(review *entity.Game) (util.Identifier, error) {
	return r.Repository.StoreGame(review)
}
