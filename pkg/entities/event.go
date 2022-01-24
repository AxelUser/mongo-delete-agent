package entities

import "time"

type Event struct {
	ClientId int64 `bson:"clientId"`
	UserId   int64 `bson:"userId"`
	Value    string
	Time     time.Time
}
