package web

import "github.com/AxelUser/mongo-delete-agent/pkg/config"

type Config struct {
	config.MongoConnection
	Port int `long:"port" default:"8080" description:"Port for web api"`
}
