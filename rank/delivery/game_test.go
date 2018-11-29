package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
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

	t.Run("should have created resource", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"name": "Game 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/games", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var game *entity.Game
		json.NewDecoder(w.Body).Decode(&game)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(game.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("shouldnt have created resource because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"name":
		  }`)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/games", strings.NewReader(payload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should get Game by ID", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{"name": "Game name"}`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/games", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var game *entity.Game
		json.NewDecoder(w.Body).Decode(&game)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(game.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodGet, "/api/v1/games/"+game.ID.String(), nil)
		router.ServeHTTP(w, req)

		var gottenGame *entity.Game
		json.NewDecoder(w.Body).Decode(&gottenGame)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("shouldnt get Game by ID because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/games/"+"adifhsghkfgiy", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should delete a Game by ID", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"name": "Game Name"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/games", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var game *entity.Game
		json.NewDecoder(w.Body).Decode(&game)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(game.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodDelete, "/api/v1/games/"+game.ID.String(), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("shouldnt delete a Game by ID because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/games/dfadfafgr", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
