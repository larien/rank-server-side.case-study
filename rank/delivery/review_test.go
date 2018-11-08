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

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controllers := repository.Review(repo)

	router := gin.Default()

	v1 := router.Group("/api/v1")

	t.Run("should create new Review handler", func(t *testing.T) {
		SetReviewEndpoints(v1, controllers)
	})

	t.Run("should set Review GET endpoint", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/review", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// TODO
	// t.Run("should set Review GET by ID endpoint", func(t *testing.T) {
	// 	id := util.NewID()

	// 	assert.NotNil(t, resp)
	// 	assert.Equal(t, id.String(), resp.review.ID.String())

	// 	assert.True(t, util.IsValidID(review.ID.String()))
	// 	assert.Equal(t, id, review.ID)
	// 	assert.Equal(t, http.StatusOK, w.Code)
	// })

	// TODO
	// t.Run("should delete Review by ID endpoint", func(t *testing.T) {
	// 	id := util.NewID()

	// 	assert.NotNil(t, resp)
	// 	assert.Equal(t, id.String(), resp.review.ID.String())

	// 	assert.True(t, util.IsValidID(review.ID.String()))
	// 	assert.Equal(t, id, review.ID)
	// 	assert.Equal(t, http.StatusOK, w.Code)
	// })

	// TODO
	// t.Run("should update Review", func(t *testing.T) {
	// 	id := util.NewID()

	// 	assert.NotNil(t, resp)
	// 	assert.Equal(t, id.String(), resp.review.ID.String())

	// 	assert.True(t, util.IsValidID(review.ID.String()))
	// 	assert.Equal(t, id, review.ID)
	// 	assert.Equal(t, http.StatusOK, w.Code)
	// })

	// TODO
	// t.Run("shouldn't be able to parse json'", func(t *testing.T) {
	// 	var jsonStr = []byte("{lalala}")

	// 	req, _ := http.NewRequest("POST", "/api/v1/review", bytes.NewBuffer(jsonStr))
	// 	req.Header.Set("Content-Type", "application/json")

	// 	client := &http.Client{}
	// 	resp, err := client.Do(req)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer resp.Body.Close()

	// 	assert.Equal(t, http.StatusInternalServerError, resp.Status)
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
