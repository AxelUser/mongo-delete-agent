package agent

import "github.com/AxelUser/mongo-delete-agent/config"

type Config struct {
	config.MongoConnection
	WCount int
	Port   int
}
