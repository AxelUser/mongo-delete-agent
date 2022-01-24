package handler

import (
	"context"
	"sync"

	"github.com/AxelUser/mongo-delete-agent/pkg/models"
)

func startWorker(ctx context.Context, qrs <-chan models.DataQuery, wg *sync.WaitGroup, op func(q models.DataQuery)) {
	go func() {
		defer wg.Done()
		for {
			select {
			case q, ok := <-qrs:
				if q == (models.DataQuery{}) && !ok {
					break
				}
				op(q)
			case <-ctx.Done():
				break
			}
		}
	}()
}
