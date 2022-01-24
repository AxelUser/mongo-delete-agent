package seed

import "github.com/AxelUser/mongo-delete-agent/pkg/config"

type Config struct {
	config.MongoConnection
	Accounts int `long:"accounts" default:"10" description:"Number of unique accounts"`
	Users    int `long:"users" default:"100" description:"Number of users for each account"`
	Events   int `long:"events" default:"100" description:"Number of events for each user"`
}
