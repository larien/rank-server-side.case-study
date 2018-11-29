package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"
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
		// endpoints.POST("", game.post)
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
