package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type ClientId int64

type UserId int64

type EventTypeId string

type DataQuery struct {
	client ClientId
	user   *UserId
}

func CreateClientQuery(id ClientId) DataQuery {
	return DataQuery{
		client: id,
		user:   nil,
	}
}

func CreateUserQuery(cId ClientId, uId UserId) DataQuery {
	return DataQuery{
		client: cId,
		user:   &uId,
	}
}

func (r DataQuery) GetFilter() bson.D {
	switch {
	case r.user == nil:
		return bson.D{
			{Key: "clientId", Value: int64(r.client)},
		}
	default:
		return bson.D{
			{Key: "clientId", Value: int64(r.client)},
			{Key: "userId", Value: int64(*r.user)},
		}
	}
}

func (r DataQuery) String() string {
	if r.user == nil {
		return fmt.Sprintf("Client: %d", r.client)
	}

	return fmt.Sprintf("Client: %d, User: %d", r.client, *r.user)
}
