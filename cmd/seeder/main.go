package main

import (
	"log"

	"github.com/AxelUser/mongo-delete-agent/config"
	"github.com/AxelUser/mongo-delete-agent/seed"
)

func main() {
	log.Println("Data-seed started")
	err := seed.Init(config.MongoConnection{
		Uri: "mongodb://localhost:27217",
		Db:  "testdb",
		Col: "testcol",
	})
	if err != nil {
		panic(err)
	}
}
