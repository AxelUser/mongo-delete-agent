package handler

import (
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func startWorker(col *mongo.Collection, reqs <-chan DeletionRequest, wg *sync.WaitGroup, op func(col *mongo.Collection, r DeletionRequest) error) {
	go func() {
		for {
			r, ok := <-reqs

			if r == (DeletionRequest{}) && !ok {
				wg.Done()
				break
			}

			err := op(col, r)
			if err != nil {
				log.Printf("failed to handle request %s: %s", r, err)
			}
		}
	}()
}
