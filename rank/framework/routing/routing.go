package routing

import (
	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/delivery"
)

// Router sets up routing for Rank app.
func Router() *gin.Engine {
	router := setup()

	endpoints(router)

	return router
}

// endpoints receives endpoints from each entity from Delivery layer.
func endpoints(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		delivery.Review(v1)
	}
}

// setup sets router with Gin framework and returns
// its default engine. It also sets up a response to the
// /hello GET request.
func setup() *gin.Engine {
	r := gin.Default()

	return r
}
