package delivery

import (
	"net/http"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"

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
	review := &ReviewHandler{
		Controller: c,
	}

	endpoints := version.Group("/review")
	{
		endpoints.GET("", review.findAll)
		endpoints.POST("", review.post)
		endpoints.GET("/:id", review.getByID)
	}
}

// findAll is the handler for GET /review endpoint.
func (r *ReviewHandler) findAll(c *gin.Context) {
	reviews, _ := r.Controller.FindAll()
	// TODO
	// reviews, err := r.Controller.FindAll()
	// if err != nil {
	// 	c.JSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{
	// 			"status":  http.StatusInternalServerError,
	// 			"message": "Failed to get reviews",
	// 			"error":   err,
	// 		})
	// 	return
	// }

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"reviews": reviews,
		})
}

// post is the handler for POST /review endpoint.
func (r *ReviewHandler) post(c *gin.Context) {
	var review entity.Review

	c.BindJSON(&review)
	// TODO
	// err := c.BindJSON(&review)
	// if err != nil {
	// 	c.JSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{
	// 			"status":  http.StatusInternalServerError,
	// 			"message": "Failed to parse json", "error": err,
	// 		})
	// 	return
	// }

	id, _ := r.Controller.Store(&review)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Review created successfully!",
			"id":      id,
		})
}

// getByID is the handler for GET /review/:id endpoint and returns desired review.
func (r *ReviewHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	if !util.IsValidID(id) {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Invalid ID",
				"error":   util.ErrInvalidID,
			})
		return
	}

	bson := util.StringToID(id)
	review, err := r.Controller.GetByID(bson)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to parse json",
				"error":   err,
			})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"review": review,
		})
}
