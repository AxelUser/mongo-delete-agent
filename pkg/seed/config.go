package seed

import "github.com/AxelUser/mongo-delete-agent/pkg/config"

type Config struct {
	config.MongoConnection
	Accounts int64 `long:"accounts" default:"10" description:"Number of unique accounts"`
	Users    int64 `long:"users" default:"10000" description:"Number of users for each"`
}
