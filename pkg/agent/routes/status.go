package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitStatus(ctx context.Context, r *gin.Engine) {
	r.GET("/status", func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	})
}
