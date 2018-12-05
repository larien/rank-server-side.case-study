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
	PublicationDate time.Time       `bson:"publication_date" json:"publication_date"`
	UpdatedAt       time.Time       `bson:"updated_at" json:"updated_at"`
	ReadingTime     string          `bson:"reading_time" json:"reading_time"`
	TextReview      string          `bson:"text_review" json:"text_review"`
	CoverImage      string          `bson:"cover_image" json:"cover_image"`
	IsPublished     bool            `bson:"is_published" json:"is_published"`
}

// Rating represents the rating entity and its atttributes.
type Rating struct {
	ID       util.Identifier `bson:"_id,omitempty" json:"id"`
	ReviewID util.Identifier `bson:"review_id,omitempty" json:"review_id"`
	// UserID   util.Identifier `bson:"user_id,omitempty" json:"user_id"`
	Rating    int       `bson:"rating" json:"rating"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
