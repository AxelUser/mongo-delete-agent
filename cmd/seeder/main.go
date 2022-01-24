package main

import (
	"context"
	"log"

	"github.com/AxelUser/mongo-delete-agent/pkg/config"
	"github.com/AxelUser/mongo-delete-agent/pkg/seed"
)

func main() {
	log.Println("Data-seed started")
	err := seed.Init(context.Background(),
		seed.Config{
			MongoConnection: config.MongoConnection{
				Uri: "mongodb://localhost:27217",
				Db:  "testdb",
				Col: "testcol",
			},
			Accounts: 10,
			Users:    1_000_000,
		})
	if err != nil {
		panic(err)
	}
}
