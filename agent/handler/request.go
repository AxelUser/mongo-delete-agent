package handler

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type ClientId int64

type UserId int64

type DeletionRequest struct {
	client ClientId
	user   *UserId
}

func CreateClientRequest(id ClientId) DeletionRequest {
	return DeletionRequest{
		client: id,
		user:   nil,
	}
}

func CreateUserRequest(cId ClientId, uId UserId) DeletionRequest {
	return DeletionRequest{
		client: cId,
		user:   &uId,
	}
}

func (r DeletionRequest) GetFilter() bson.D {
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

func (r DeletionRequest) String() string {
	if r.user == nil {
		return fmt.Sprintf("Client: %d", r.client)
	}

	return fmt.Sprintf("Client: %d, User: %d", r.client, *r.user)
}
