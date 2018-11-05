package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouter_GET_Review(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/review/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, Rank!", w.Body.String())
}
