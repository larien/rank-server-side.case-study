package delivery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEndpoint_GET_Review(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		Review(v1)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/review/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
