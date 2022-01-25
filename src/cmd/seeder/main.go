package main

import (
	"context"
	"log"

	"github.com/AxelUser/mongo-delete-agent/src/seed"
	"github.com/jessevdk/go-flags"
)

func main() {
	var c seed.Config
	_, err := flags.Parse(&c)
	if err != nil {
		panic(err)
	}

	log.Println("Data-seed started")
	err = seed.Init(context.Background(), c)
	if err != nil {
		panic(err)
	}
}
