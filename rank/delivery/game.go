package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

// Game contains injected interface from Controller layer.
type Game struct {
	Controller controller.GameController
}

// SetGameEndpoints sets endpoints for Game entity.
func SetGameEndpoints(version *gin.RouterGroup, c controller.GameController) {
	game := &Game{
		Controller: c,
	}

	endpoints := version.Group("/games")
	{
		endpoints.GET("", game.findAll)
		// endpoints.GET("/:id", game.getByID)
		endpoints.POST("", game.post)
		// endpoints.PATCH("", game.patch)
		// endpoints.DELETE("/:id", game.deleteByID)
	}
}

// findAll handles GET /games requests and returns all Games from database.
func (g *Game) findAll(c *gin.Context) {
	games, _ := g.Controller.FindAllGames()

	c.JSON(
		http.StatusOK,
		games,
	)
}

// post handles POST /games requests on creating a new Game.
func (g *Game) post(c *gin.Context) {
	var game entity.Game

	if err := c.BindJSON(&game); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to parse json",
				"error":   err,
			})
		return
	}

	id, _ := g.Controller.StoreGame(&game)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Review created successfully!",
			"id":      id,
		})
}
