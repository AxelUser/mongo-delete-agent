package entities

import "time"

type Event struct {
	ClientId int64  `bson:"clientId"`
	UserId   int64  `bson:"userId"`
	TypeId   string `bson:"typeId"`
	Props    map[string]string
	Time     time.Time
}
