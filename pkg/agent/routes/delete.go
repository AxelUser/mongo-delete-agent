package routes

import (
	"context"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/pkg/agent/handler"
	"github.com/AxelUser/mongo-delete-agent/pkg/agent/requests"
	"github.com/AxelUser/mongo-delete-agent/pkg/storage"
	"github.com/gin-gonic/gin"
)

func InitDelete(ctx context.Context, r *gin.Engine, repo *storage.EventsRepository, h *handler.Handler) {
	delete := r.Group("/delete")

	delete.POST("/:clientId", func(c *gin.Context) {
		var req requests.ClientReq

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.Delete(ctx, req.Query())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})

	delete.POST("/:clientId/:userId", func(c *gin.Context) {
		var req requests.UserReq

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.Delete(ctx, req.Query())
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})
}
