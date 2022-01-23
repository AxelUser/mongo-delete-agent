package main

import (
	"log"

	"github.com/AxelUser/mongo-delete-agent/config"
	"github.com/AxelUser/mongo-delete-agent/seed"
)

func main() {
	log.Println("Data-seed started")
	err := seed.Init(
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
