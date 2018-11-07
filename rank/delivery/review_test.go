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

func TestEndpoint_Review(t *testing.T) {

	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.NewMongoConnection(pool, config.MONGODB_DATABASE)

	controllers := repository.Review(repo)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	t.Run("should create new Review handler", func(t *testing.T) {
		NewReviewHandler(v1, controllers)
	})

	t.Run("should set Review endpoints", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/review", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// TODO
	// t.Run("shouldn't be able to parse json'", func(t *testing.T) {
	// 	w := httptest.NewRecorder()

	// 	form := url.Values{}
	// 	form.Add("username", "username")
	// 	req := httptest.NewRequest("POST", "/api/v1/review", strings.NewReader(form.Encode()))
	// 	req.Form = form
	// 	router.ServeHTTP(w, req)

	// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// })

	t.Run("should have created resource", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title": "Title 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/review", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
