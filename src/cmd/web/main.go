package main

import (
	"context"

	"github.com/AxelUser/mongo-delete-agent/src/web"
	"github.com/jessevdk/go-flags"
)

func main() {
	var c web.Config
	_, err := flags.Parse(&c)
	if err != nil {
		panic(err)
	}

	web.Start(context.Background(), c)
}
