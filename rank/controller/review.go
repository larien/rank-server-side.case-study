package controller

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review contains the injected Review interface from Repository layer.
type Review struct {
	Repository repository.Review
}

// ReviewController contains methods that must be implemented by the injected layer.
type ReviewController interface {
	DeleteReviewByID(util.Identifier) error
	FindAllReviews() ([]*entity.Review, error)
	FindAllUnpublishedReviews() ([]*entity.Review, error)
	GetReviewByID(util.Identifier) (*entity.Review, error)
	StoreReview(*entity.Review) (util.Identifier, error)
	UpdateReview(*entity.Review) error
}

// newReviewController creates a new Review Controller.
func newReviewController(m *repository.MongoDB) *Review {
	return &Review{
		Repository: m,
	}
}

// DeleteReviewByID requests the Repository layer for a Review to be deleted from the database by its ID.
func (r *Review) DeleteReviewByID(id util.Identifier) error {
	return r.Repository.DeleteReviewByID(id)
}

// FindAllReviews requests the Repository layer to return all published Reviews from database.
func (r *Review) FindAllReviews() ([]*entity.Review, error) {
	return r.Repository.FindAllReviews()
}

// FindAllUnpublishedReviews requests the Repository layer to return all unpublished Reviews from database.
func (r *Review) FindAllUnpublishedReviews() ([]*entity.Review, error) {
	return r.Repository.FindAllUnpublishedReviews()
}

// GetReviewByID requests the Repository layer for a certain Review by its ID.
func (r *Review) GetReviewByID(id util.Identifier) (*entity.Review, error) {
	return r.Repository.GetReviewByID(id)
}

// StoreReview requests the Repository layer for the insertion of a new Review in the database.
func (r *Review) StoreReview(review *entity.Review) (util.Identifier, error) {
	return r.Repository.StoreReview(review)
}

// UpdateReview requests the Repository layer for a Review to be updated in the database.
func (r *Review) UpdateReview(review *entity.Review) error {
	return r.Repository.UpdateReview(review)
}

// RateReview requests the Repository layer for the insertion of a new Rating in the database.
func (r *Review) RateReview(rating *entity.Rating) (util.Identifier, error) {
	return r.Repository.RateReview(rating)
}

// FindRatingsByReview returns all Ratings from a certain Review.
func (r *Review) FindRatingsByReview(reviewID util.Identifier) ([]*entity.Rating, error) {
	return r.Repository.FindRatingsByReview(reviewID)
}

// GetAverageRating calculates all the ratings received by a review and returns its average rating.
func (r *Review) GetAverageRating(reviewID util.Identifier) (int, error) {
	ratings, _ := r.FindRatingsByReview(reviewID)

	var sum int
	for _, v := range ratings {
		sum += v.Rate
	}

	return sum / len(ratings), nil
}
