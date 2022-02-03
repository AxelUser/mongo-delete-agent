package routes

import (
	"context"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/src/agent/requests"
	"github.com/AxelUser/mongo-delete-agent/src/storage"
	"github.com/gin-gonic/gin"
)

func InitStatus(ctx context.Context, r *gin.Engine, h *storage.AdminHelper) {
	r.GET("/status", func(c *gin.Context) {
		var req requests.JobReq
		err := c.ShouldBindHeader(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		run, err := h.IsCommandRunning(ctx, req.JobId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if run {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusNotFound)
		}
	})
}
