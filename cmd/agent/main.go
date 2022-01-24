package main

import (
	"context"
	"log"

	"github.com/AxelUser/mongo-delete-agent/pkg/agent"
	"github.com/AxelUser/mongo-delete-agent/pkg/config"
)

func main() {
	log.Println("Mongo deletion agent started")
	err := agent.Start(context.Background(), agent.Config{
		MongoConnection: config.MongoConnection{
			Uri: "mongodb://localhost:27217",
			Db:  "testdb",
			Col: "testcol",
		},
		WCount: 2,
		Port:   8080,
	})

	if err != nil {
		panic(err)
	}
}
