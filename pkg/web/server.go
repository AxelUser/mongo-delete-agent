package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AxelUser/mongo-delete-agent/pkg/models"
	"github.com/AxelUser/mongo-delete-agent/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, c Config) error {
	repo, err := storage.CreateEventsRepository(ctx, c.MongoConnection)
	if err != nil {
		return fmt.Errorf("failed to start API: %w", err)
	}

	router := gin.Default()

	router.GET("/events/stats", func(c *gin.Context) {
		m, err := repo.CountByAccount(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, m)
	})

	router.GET("/events/:clientId/:userId/:typeId", func(c *gin.Context) {
		r := struct {
			ClientId int64  `uri:"clientId" binding:"required"`
			UserId   int64  `uri:"userId" binding:"required"`
			TypeId   string `uri:"typeId" binding:"required"`
			Skip     int64  `form:"skip"`
			Take     int64  `form:"take"`
		}{Take: 100}

		c.BindUri(&r)
		c.BindQuery(&r)

		l, err := repo.Get(ctx, models.ClientId(r.ClientId), models.UserId(r.UserId), models.EventTypeId(r.TypeId), r.Skip, r.Take)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(l) == 0 {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, l)
		}
	})

	router.Run(fmt.Sprintf(":%d", c.Port))

	return nil
}
