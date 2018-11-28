package routing

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/delivery"
)

// Router sets up routing for Rank app.
func Router(controllers *controller.Controllers) *gin.Engine {
	router := setup()

	router.Use(cors.Default())

	endpoints(router, controllers)

	return router
}

// endpoints receives endpoints from each entity from Delivery layer.
func endpoints(router *gin.Engine, controllers *controller.Controllers) {
	v1 := router.Group("/api/v1")
	{
		delivery.SetReviewEndpoints(v1, controllers.Review)
	}
}

// setup sets router with Gin framework and returns
// its default engine. It also sets up a response to the
// /hello GET request.
func setup() *gin.Engine {
	r := gin.Default()

	return r
}
