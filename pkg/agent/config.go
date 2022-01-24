package agent

import "github.com/AxelUser/mongo-delete-agent/pkg/config"

type Config struct {
	config.MongoConnection
	WCount int `long:"workers" default:"10" description:"Amount of workers that handle deletion"`
	Port   int `long:"port" default:"8080" description:"Port for agent"`
}
