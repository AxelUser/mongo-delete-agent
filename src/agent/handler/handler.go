package handler

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/AxelUser/mongo-delete-agent/src/models"
	"github.com/AxelUser/mongo-delete-agent/src/storage"
)

type Handler struct {
	jobs   chan<- job
	dones  *sync.WaitGroup
	wCount int
}

func Create(ctx context.Context, repo storage.EventsRepository, wCount int) (*Handler, error) {
	jobs := make(chan job)
	var dones sync.WaitGroup
	dones.Add(wCount)
	for i := 0; i < wCount; i++ {
		startWorker(ctx, jobs, &dones, func(j job) {
			delCount, err := repo.Delete(context.Background(), j.id, j.query)
			if err != nil {
				log.Printf("Failed to delete events: %s", err)
			}

			log.Printf("Deleted %d events for request '%s'", delCount, j.query)
		})
	}

	return &Handler{
		jobs:   jobs,
		dones:  &dones,
		wCount: wCount,
	}, nil
}

func (h Handler) Delete(ctx context.Context, jobId int64, q models.DataQuery) error {
	select {
	case <-ctx.Done():
		return nil
	case h.jobs <- job{
		id:    jobId,
		query: q,
	}:
		return nil
	case <-time.After(15 * time.Second):
		return errors.New("failed to schedule delete, timeout 15 seconds")
	}
}
