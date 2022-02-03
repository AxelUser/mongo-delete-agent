package routes

import (
	"context"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/src/agent/handler"
	"github.com/AxelUser/mongo-delete-agent/src/agent/requests"
	"github.com/AxelUser/mongo-delete-agent/src/storage"
	"github.com/gin-gonic/gin"
)

func InitDelete(ctx context.Context, r *gin.Engine, repo *storage.EventsRepository, h *handler.Handler) {
	delete := r.Group("/delete")

	delete.POST("/:clientId", func(c *gin.Context) {
		var job requests.JobReq
		err := c.ShouldBindHeader(&job)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter requests.ClientReq
		err = c.ShouldBindUri(&filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.Delete(ctx, job.JobId, filter.Query())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})

	delete.POST("/:clientId/:userId", func(c *gin.Context) {
		var job requests.JobReq
		err := c.ShouldBindHeader(&job)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter requests.UserReq
		err = c.ShouldBindUri(&filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.Delete(ctx, job.JobId, filter.Query())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})
}
