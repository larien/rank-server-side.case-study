package repository

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/middlewares/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

func TestFindAllReviews(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have returned all reviews", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		m.StoreReview(r1)
		reviews, err := m.FindAllReviews()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
	})

	t.Run("should have returned error", func(t *testing.T) {
		m = New(pool, "otherdatabase")
		reviews, err := m.FindAllReviews()
		assert.NotNil(t, err)
		assert.Nil(t, reviews)
	})
}

func TestFindAllUnpublishedReviews(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have returned all unpublished reviews", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		m.StoreReview(r1)
		reviews, err := m.FindAllUnpublishedReviews()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
		assert.False(t, reviews[0].IsPublished)
	})

	t.Run("should have returned error", func(t *testing.T) {
		m = New(pool, "otherdatabase")
		reviews, err := m.FindAllUnpublishedReviews()
		assert.NotNil(t, err)
		assert.Nil(t, reviews)
	})
}

func TestStoreReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have inserted a new review", func(t *testing.T) {
		// TODO - change to RemoveAll function from Repository layer
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		// TODO
		id, _ := m.StoreReview(r1)

		reviews, errFindAll := m.FindAllReviews()

		assert.Equal(t, 1, len(reviews))
		assert.Nil(t, err)
		assert.Nil(t, errFindAll)
		assert.NotNil(t, id)
		assert.Equal(t, true, util.IsValidID(id.String()))
	})
}

func TestFindReviewByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should find certain Review by stored ID", func(t *testing.T) {

		r1 := &entity.Review{
			Title: "Title Test",
		}

		// TODO
		id, _ := m.StoreReview(r1)

		review, err := m.GetReviewByID(id)
		assert.Equal(t, "Title Test", review.Title)
		assert.Nil(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, review.ID)
		assert.Equal(t, true, util.IsValidID(id.String()))
		assert.Equal(t, true, util.IsValidID(review.ID.String()))
	})
}

func TestDeleteReviewByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should delete certain Review by stored ID", func(t *testing.T) {

		r1 := &entity.Review{
			Title: "Title Test",
		}

		// TODO
		id, _ := m.StoreReview(r1)

		review, errGetByID := m.GetReviewByID(id)
		assert.Equal(t, id, review.ID)
		assert.Nil(t, errGetByID)

		err := m.DeleteReviewByID(id)
		assert.Nil(t, err)

		review, errGetByID2 := m.GetReviewByID(id)
		assert.Nil(t, review)
		assert.Nil(t, errGetByID2)
	})
}

func TestUpdateReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have updated a new review", func(t *testing.T) {
		// TODO - change to RemoveAll function from Repository layer
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

		r1 := &entity.Review{
			Title: "Title 1",
		}

		// TODO
		id, err := m.StoreReview(r1)
		assert.Nil(t, err)

		review, errGetByID := m.GetReviewByID(id)
		assert.Nil(t, errGetByID)
		assert.Equal(t, "Title 1", review.Title)

		review.Title = "Different title"
		errUpdate := m.UpdateReview(review)
		assert.Nil(t, errUpdate)

		updatedReview, errGetByID2 := m.GetReviewByID(id)

		assert.Nil(t, errGetByID2)
		assert.Equal(t, "Different title", updatedReview.Title)
	})
}
