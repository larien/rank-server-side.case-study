package main

import (
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {
	Rank()
}

// Rank starts the routine for Rank's app.
func Rank() {
	router := setupRouter()
	router.Run(port)
}

// setupRouter sets router with Gin framework and returns
// its default engine. It also sets up a response to the
// /hello GET request.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Rank!")
	})

	return r
}
