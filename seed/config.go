package seed

import "github.com/AxelUser/mongo-delete-agent/config"

type Config struct {
	config.MongoConnection
	Accounts int64
	Users    int64
}
