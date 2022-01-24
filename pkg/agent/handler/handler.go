package handler

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/AxelUser/mongo-delete-agent/pkg/models"
	"github.com/AxelUser/mongo-delete-agent/pkg/storage"
)

type Handler struct {
	reqs   chan<- models.DataQuery
	dones  *sync.WaitGroup
	wCount int
}

func Create(ctx context.Context, repo storage.EventsRepository, wCount int) (*Handler, error) {
	reqs := make(chan models.DataQuery)
	var dones sync.WaitGroup
	dones.Add(wCount)
	for i := 0; i < wCount; i++ {
		startWorker(ctx, reqs, &dones, func(q models.DataQuery) {
			delCount, err := repo.Delete(context.Background(), q)
			if err != nil {
				log.Printf("Failed to delete events: %s", err)
			}

			log.Printf("Deleted %d events for request '%s'", delCount, q)
		})
	}

	return &Handler{
		reqs:   reqs,
		dones:  &dones,
		wCount: wCount,
	}, nil
}

func (h Handler) Delete(ctx context.Context, r models.DataQuery) error {
	select {
	case <-ctx.Done():
		return nil
	case h.reqs <- r:
		return nil
	case <-time.After(15 * time.Second):
		return errors.New("failed to schedule delete, timeout 15 seconds")
	}
}
