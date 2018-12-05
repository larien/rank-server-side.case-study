package repository

import (
	"time"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/middlewares/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	"gopkg.in/mgo.v2/bson"
)

// Review defines the methods must be implemented by injected layer.
type Review interface {
	DeleteReviewByID(util.Identifier) error
	FindAllReviews() ([]*entity.Review, error)
	FindAllUnpublishedReviews() ([]*entity.Review, error)
	GetReviewByID(util.Identifier) (*entity.Review, error)
	StoreReview(*entity.Review) (util.Identifier, error)
	UpdateReview(*entity.Review) error
	RateReview(*entity.Rating) (util.Identifier, error)
	FindRatingsByReview(reviewID util.Identifier) ([]*entity.Rating, error)
}

// DeleteReviewByID deletes a Review by its ID.
func (m *MongoDB) DeleteReviewByID(id util.Identifier) error {
	return m.pool.Session(nil).DB(m.db).C(config.REVIEW_COLLECTION).RemoveId(id)
}

// FindAllReviews returns all published Reviews from the database sorted by ID.
func (m *MongoDB) FindAllReviews() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.REVIEW_COLLECTION)
	if err := collection.Find(bson.M{"is_published": true}).Sort("id").All(&reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

// FindAllUnpublishedReviews returns all unpublished Reviews from the database sorted by ID.
func (m *MongoDB) FindAllUnpublishedReviews() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.REVIEW_COLLECTION)
	if err := collection.Find(bson.M{"is_published": false}).Sort("id").All(&reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetReviewByID finds a Review by its ID.
func (m *MongoDB) GetReviewByID(id util.Identifier) (*entity.Review, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	var review *entity.Review

	coll.FindId(id).One(&review)

	return review, nil
}

// StoreReview inserts a new Review in the database.
func (m *MongoDB) StoreReview(review *entity.Review) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	review.ID = util.NewID() // TODO - improve function
	review.UpdatedAt = time.Now()
	review.IsPublished = false

	coll.Insert(review)

	return review.ID, nil
}

// UpdateReview updates an existing Review changing is_published and updated_at fields in the database.
func (m *MongoDB) UpdateReview(review *entity.Review) error {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	err := coll.UpdateId(review.ID, bson.M{"$set": bson.M{"is_published": true, "updated_at": time.Now()}}) // TODO - avoid null Reviews

	return err
}

// RateReview inserts a new Rating in the database.
func (m *MongoDB) RateReview(rating *entity.Rating) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.RATING_COLLECTION)

	rating.ID = util.NewID() // TODO - improve function
	rating.UpdatedAt = time.Now()

	coll.Insert(rating)

	return rating.ID, nil
}

// FindAllRatings returns all Ratings from the database sorted by ID.
func (m *MongoDB) FindAllRatings() ([]*entity.Rating, error) {
	var ratings []*entity.Rating

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.RATING_COLLECTION)
	if err := collection.Find(nil).Sort("id").All(&ratings); err != nil {
		return nil, err
	}

	return ratings, nil
}

// FindRatingsByReview returns all Ratings of a certain Review from the database sorted by ID.
func (m *MongoDB) FindRatingsByReview(reviewID util.Identifier) ([]*entity.Rating, error) {
	var ratings []*entity.Rating

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.RATING_COLLECTION)
	if err := collection.Find(bson.M{"review_id": reviewID}).Sort("id").All(&ratings); err != nil {
		return nil, err
	}

	return ratings, nil
}
