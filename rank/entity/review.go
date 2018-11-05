package entity

import (
	"time"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review represents the review entity and its attributes.
type Review struct {
	ID              util.Identifier `bson:"_id,omitempty" json:"id"`
	Title           string          `bson:"title" json:"title"`
	AuthorNickname  string          `bson:"author_nickname" json:"author_nickname"`
	AverageRating   int             `bson:"average_rating" json:"average_rating"`
	PublicationDate time.Time       `bson:"publication_date" json:"publication_date"`
	ReadingTime     int             `bson:"reading_time" json:"reading_time"`
	TextReview      string          `bson:"text_review" json:"text_review"`
	CoverImage      string          `bson:"cover_image" json:"cover_image"`
}
