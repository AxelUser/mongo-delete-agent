package requests

import "github.com/AxelUser/mongo-delete-agent/pkg/models"

type UserReq struct {
	ClientId int64 `uri:"clientId" binding:"required"`
	UserId   int64 `uri:"userId" binding:"required"`
}

func (r UserReq) Query() models.DataQuery {
	return models.CreateUserQuery(models.ClientId(r.ClientId), models.UserId(r.UserId))
}
