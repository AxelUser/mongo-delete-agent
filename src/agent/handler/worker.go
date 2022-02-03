package handler

import (
	"context"
	"sync"
)

func startWorker(ctx context.Context, jobs <-chan job, wg *sync.WaitGroup, op func(j job)) {
	go func() {
		defer wg.Done()
		for {
			select {
			case j, ok := <-jobs:
				if !ok {
					return
				}
				op(j)
			case <-ctx.Done():
				return
			}
		}
	}()
}
