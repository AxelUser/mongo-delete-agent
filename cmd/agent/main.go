package main

import (
	"log"

	"github.com/AxelUser/mongo-delete-agent/agent"
	"github.com/AxelUser/mongo-delete-agent/config"
)

func main() {
	log.Println("Mongo deletion agent started")
	err := agent.Start(agent.Config{
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
