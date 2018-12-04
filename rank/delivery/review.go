package delivery

import (
	"net/http"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"

	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"

	"github.com/gin-gonic/gin"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
)

// Review contains injected interface from Controller layer.
type Review struct {
	Controller controller.ReviewController
}

// SetReviewEndpoints sets endpoints for Review entity.
func SetReviewEndpoints(version *gin.RouterGroup, c controller.ReviewController) {
	review := &Review{
		Controller: c,
	}

	endpoints := version.Group("/reviews")
	{
		endpoints.GET("", review.findAll)
		endpoints.GET("/:id", review.getByID)
		endpoints.POST("", review.post)
		endpoints.PATCH("", review.patch)
		endpoints.DELETE("/:id", review.deleteByID)
	}
}

// findAll handles GET /review requests and returns all Reviews from database.
func (r *Review) findAll(c *gin.Context) {
	reviews, _ := r.Controller.FindAllReviews()

	c.JSON(
		http.StatusOK,
		reviews,
	)
}

// post handles POST /review requests on creating a new Review.
func (r *Review) post(c *gin.Context) {

	if !authorizate(c.Request.Header.Get("Authorization")) {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		return
	}

	var review entity.Review

	if err := c.BindJSON(&review); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to parse json",
				"error":   err,
			})
		return
	}

	id, _ := r.Controller.StoreReview(&review)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Review created successfully!",
			"id":      id,
		})
}

// getByID handles GET /review/:id requests and returns desired Review by its ID.
func (r *Review) getByID(c *gin.Context) {
	id := c.Param("id")
	if !util.IsValidID(id) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid ID",
				"error":   util.ErrInvalidID,
			})
		return
	}

	bson := util.StringToID(id)
	review, _ := r.Controller.GetReviewByID(bson)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"review": review,
		})
}

// deleteByID handles DELETE /review/:id requests and deletes desired Review by its ID.
func (r *Review) deleteByID(c *gin.Context) {
	if !authorizate(c.Request.Header.Get("Authorization")) {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		return
	}

	id := c.Param("id")
	if !util.IsValidID(id) {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid ID",
				"error":   util.ErrInvalidID,
			})
		return
	}

	bson := util.StringToID(id)
	r.Controller.DeleteReviewByID(bson)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
		})
}

// patch handles PATCH /review endpoint and updates an existing Review.
func (r *Review) patch(c *gin.Context) {
	if !authorizate(c.Request.Header.Get("Authorization")) {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
		return
	}

	var review entity.Review

	if err := c.BindJSON(&review); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Failed to parse json",
				"error":   err,
			})
		return
	}

	r.Controller.UpdateReview(&review)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Review updated successfully!",
		})
}
