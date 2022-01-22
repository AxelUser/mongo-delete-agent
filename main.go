package main

import (
	"fmt"

	"github.com/AxelUser/mongo-delete-agent/seed"
)

func main() {
	fmt.Println("Mongo delete agent")
	err := seed.Init("mongodb://localhost:27217", "testdb", "testcol")
	if err != nil {
		panic(err)
	}
}
