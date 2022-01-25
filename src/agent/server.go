package agent

import (
	"context"
	"fmt"

	"github.com/AxelUser/mongo-delete-agent/src/agent/handler"
	"github.com/AxelUser/mongo-delete-agent/src/agent/routes"
	"github.com/AxelUser/mongo-delete-agent/src/storage"
	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, conf Config) error {
	repo, err := storage.CreateEventsRepository(ctx, conf.MongoConnection)
	if err != nil {
		return fmt.Errorf("failed to start agent: %w", err)
	}

	h, err := handler.Create(ctx, *repo, conf.WCount)
	if err != nil {
		return fmt.Errorf("failed to start agent: %w", err)
	}

	router := gin.Default()

	routes.InitDelete(ctx, router, repo, h)
	routes.InitExists(ctx, router, repo, h)
	routes.InitStatus(ctx, router)

	router.Run(fmt.Sprintf(":%d", conf.Port))

	return nil
}
