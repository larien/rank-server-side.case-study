package delivery

import (
	"fmt"
	"net/http"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"

	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

type ReviewHandler struct {
	Controller controller.ReviewController
}

// NewReviewHandler prepares endpoints for Review entity.
func NewReviewHandler(version *gin.RouterGroup, r controller.ReviewController) {
	handler := &ReviewHandler{
		Controller: r,
	}
	endpoints := version.Group("/review")
	{
		endpoints.GET("/", handler.fetchAllReviews)
		endpoints.POST("/", handler.postReview)
	}
}

func (r *ReviewHandler) fetchAllReviews(c *gin.Context) {
	c.String(200, "Hello, Rank!")

	reviews, err := r.Controller.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to get reviews", "error": err})
	}
	fmt.Printf("Reviews: %+v", reviews)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "reviews": reviews})
}

func (r *ReviewHandler) postReview(c *gin.Context) {
	review := entity.Review{Title: c.PostForm("title")}
	fmt.Printf("Review recebida: %+v", review)

	// chama função no controller
	id, err := r.Controller.Store(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to create review", "error": err})
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Review created successfully!", "id": id})
}
