package delivery

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	mgo "gopkg.in/mgo.v2"
)

func TestGameEndpoints(t *testing.T) {

	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := repository.Game(repo)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	t.Run("should create new Game handler", func(t *testing.T) {
		SetGameEndpoints(v1, controller)
	})

	t.Run("should set Game GET endpoint", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/games", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
