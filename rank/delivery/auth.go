package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const expectedToken = "340CT"

// authorizate verifies token and authorizates request.
func authorizate(token string) bool {
	return token == expectedToken
}

// signIn verifies token and authorizates request.
func signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"token":  expectedToken,
			})
	}
}
