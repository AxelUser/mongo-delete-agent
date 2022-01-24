package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AxelUser/mongo-delete-agent/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	col    *mongo.Collection
	reqs   chan<- DeletionRequest
	dones  *sync.WaitGroup
	wCount int
}

func Create(con config.MongoConnection, wCount int) (*Handler, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(con.Uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	col := client.Database(con.Db).Collection(con.Col)

	reqs := make(chan DeletionRequest)
	var dones sync.WaitGroup
	dones.Add(wCount)
	for i := 0; i < wCount; i++ {
		startWorker(col, reqs, &dones, simpleDelete)
	}

	return &Handler{
		col:    col,
		reqs:   reqs,
		dones:  &dones,
		wCount: wCount,
	}, nil
}

func simpleDelete(col *mongo.Collection, r DeletionRequest) error {
	res, err := col.DeleteMany(context.Background(), r.GetFilter())
	if err != nil {
		return err
	}

	log.Printf("Deleted %d events for request %s", res.DeletedCount, r)
	return nil
}

func (h Handler) Delete(r DeletionRequest) error {
	select {
	case h.reqs <- r:
		return nil
	case <-time.After(15 * time.Second):
		return errors.New("failed to schedule delete, timeout 15 seconds")
	}

}
