package agent

import (
	"fmt"
	"strconv"

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
		if client, err := strconv.Atoi(c.Param("clientId")); err == nil {
			h.Delete(handler.CreateClientRequest(handler.ClientId(client)))
			c.Status(202)
		} else {
			c.Status(400)
		}
	})

	router.POST("/delete/:clientId/:userId", func(c *gin.Context) {
		client, err := strconv.Atoi(c.Param("clientId"))
		if err != nil {
			c.Status(500)
			return
		}

		user, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.Status(500)
			return
		}
		h.Delete(handler.CreateUserRequest(handler.ClientId(client), handler.UserId(user)))
		c.Status(202)
	})

	router.Run(fmt.Sprintf(":%d", conf.Port))

	return nil
}
