package handler

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func startWorker(col *mongo.Collection, reqs <-chan DeletionRequest, dones chan<- bool, op func(col *mongo.Collection, r DeletionRequest) error) {
	go func() {
		for {
			r, closed := <-reqs

			if r == (DeletionRequest{}) && closed {
				dones <- true
				break
			}

			err := op(col, r)
			if err != nil {
				log.Printf("failed to handle request %s: %s", r, err)
			}
		}
	}()
}
