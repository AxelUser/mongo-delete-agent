package routes

import (
	"context"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/src/agent/handler"
	"github.com/AxelUser/mongo-delete-agent/src/agent/requests"
	"github.com/AxelUser/mongo-delete-agent/src/models"
	"github.com/AxelUser/mongo-delete-agent/src/storage"
	"github.com/gin-gonic/gin"
)

func InitExists(ctx context.Context, r *gin.Engine, repo *storage.EventsRepository, h *handler.Handler) {
	exists := r.Group("/exists")

	exists.GET("/:clientId", func(c *gin.Context) {
		var req requests.ClientReq

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		handleExist(ctx, repo, c, models.CreateClientQuery(models.ClientId(req.ClientId)))
	})

	exists.GET("/:clientId/:userId", func(c *gin.Context) {
		var req requests.UserReq

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		handleExist(ctx, repo, c, models.CreateClientQuery(models.ClientId(req.ClientId)))
	})
}

func handleExist(ctx context.Context, repo *storage.EventsRepository, c *gin.Context, q models.DataQuery) {
	exists, err := repo.Exists(ctx, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotFound)
	}
}
