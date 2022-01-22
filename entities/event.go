package entities

import "time"

type Event struct {
	ClientId int64
	UserId   int64
	Value    string
	Time     time.Time
}
