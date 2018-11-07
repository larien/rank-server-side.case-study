package delivery

import (
	"net/http"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"

	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

// ReviewHandler contains injected interface from Controller layer.
type ReviewHandler struct {
	Controller controller.ReviewController
}

// NewReviewHandler prepares endpoints for Review entity.
func NewReviewHandler(version *gin.RouterGroup, c controller.ReviewController) {
	handler := &ReviewHandler{
		Controller: c,
	}

	endpoints := version.Group("/review")
	{
		endpoints.GET("", handler.fetchAllReviews)
		endpoints.POST("", handler.postReview)
	}
}

// fetchAllReviews is the handler for GET /review endpoint.
func (r *ReviewHandler) fetchAllReviews(c *gin.Context) {
	reviews, err := r.Controller.FindAll()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to get reviews",
				"error":   err,
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"reviews": reviews,
		})
}

// postReview is the handler for POST /review endpoint.
func (r *ReviewHandler) postReview(c *gin.Context) {
	var review entity.Review

	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to parse json", "error": err,
			})
		return
	}

	id := r.Controller.Store(&review)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Review created successfully!",
			"id":      id,
		})
}
