package main

import (
	"context"
	"log"

	"github.com/AxelUser/mongo-delete-agent/src/agent"
	"github.com/jessevdk/go-flags"
)

func main() {
	var c agent.Config
	_, err := flags.Parse(&c)
	if err != nil {
		panic(err)
	}

	log.Println("Mongo deletion agent started")
	err = agent.Start(context.Background(), c)

	if err != nil {
		panic(err)
	}
}
