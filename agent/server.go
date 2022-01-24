package agent

import (
	"fmt"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/agent/handler"
	"github.com/gin-gonic/gin"
)

func Start(conf Config) error {

	h, err := handler.Create(conf.MongoConnection, conf.WCount)
	if err != nil {
		return fmt.Errorf("failed to start agent: %w", err)
	}

	router := gin.Default()

	router.POST("/delete/:clientId", func(c *gin.Context) {
		req := struct {
			ClientId int64 `uri:"clientId" binding:"required"`
		}{}

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.Delete(handler.CreateClientRequest(handler.ClientId(req.ClientId)))
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})

	router.POST("/delete/:clientId/:userId", func(c *gin.Context) {
		req := struct {
			ClientId int64 `uri:"clientId" binding:"required"`
			UserId   int64 `uri:"userId" binding:"required"`
		}{}

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.Delete(handler.CreateUserRequest(handler.ClientId(req.ClientId), handler.UserId(req.UserId)))
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusAccepted)
	})

	router.Run(fmt.Sprintf(":%d", conf.Port))

	return nil
}
