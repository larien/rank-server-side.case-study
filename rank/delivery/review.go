package delivery

import (
	"github.com/gin-gonic/gin"
)

// Review prepares endpoints for Review entity.
func Review(version *gin.RouterGroup) {
	endpoints := version.Group("/review")
	{
		endpoints.GET("/", test)
	}
}

func test(c *gin.Context) {
	c.String(200, "Hello, Rank!")
}
