package web

import "github.com/AxelUser/mongo-delete-agent/src/config"

type Config struct {
	config.MongoConnection
	Port int `long:"port" default:"80" description:"Port for web api"`
}
