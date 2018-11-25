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
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/reviews", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should have created resource", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title": "Title 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("shouldnt have created resource because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title":
		  }`)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should get Review by ID", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{"title": "Title 1"}`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodGet, "/api/v1/reviews/"+review.ID.String(), nil)
		router.ServeHTTP(w, req)

		var gottenReview *entity.Review
		json.NewDecoder(w.Body).Decode(&gottenReview)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("shouldnt get Review by ID because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/reviews/"+"adifhsghkfgiy", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should delete a Review by ID", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title": "Title 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)

		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)
		fmt.Print(review)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodDelete, "/api/v1/reviews/"+review.ID.String(), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("shouldnt delete a Review by ID because of bad syntax", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/reviews/dfadfafgr", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should update a Review", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title": "Title 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)
		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		newPayload := fmt.Sprintf(`{
			"id": "` + review.ID.String() + `",
			"title": "New Title"
		  }`)
		req, err = http.NewRequest(http.MethodPatch, "/api/v1/reviews", strings.NewReader(newPayload))
		router.ServeHTTP(w, req)

		var response *entity.Review
		json.NewDecoder(w.Body).Decode(&response)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("shouldnt update a Review because of wrong syntax", func(t *testing.T) {
		w := httptest.NewRecorder()

		payload := fmt.Sprintf(`{
			"title": "Title 1"
		  }`)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/reviews", strings.NewReader(payload))
		router.ServeHTTP(w, req)
		var review *entity.Review
		json.NewDecoder(w.Body).Decode(&review)
		assert.Nil(t, err)
		assert.True(t, util.IsValidID(review.ID.String()))
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		newPayload := fmt.Sprintf(`{
			"id": "` + review.ID.String() + `",
			"title": 
		  }`)
		req, err = http.NewRequest(http.MethodPatch, "/api/v1/reviews", strings.NewReader(newPayload))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
