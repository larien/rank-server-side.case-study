package entity

import (
	"time"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review represents the review entity and its attributes.
type Review struct {
	ID              bson.Identifier `bson:"_id,omitempty"`
	Title           string
	AuthorNickname  string
	AverageRating   int
	PublicationDate time.Time
	ReadingTime     int
	TextReview      string
	ThumbnailURL    string
}
