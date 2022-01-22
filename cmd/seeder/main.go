package main

import (
	"log"

	"github.com/AxelUser/mongo-delete-agent/seed"
)

func main() {
	log.Println("Mongo delete seeder started")
	err := seed.Init("mongodb://localhost:27217", "testdb", "testcol")
	if err != nil {
		panic(err)
	}
}
