package requests

import "github.com/AxelUser/mongo-delete-agent/pkg/models"

type ClientReq struct {
	ClientId int64 `uri:"clientId" binding:"required"`
}

func (r ClientReq) Query() models.DataQuery {
	return models.CreateClientQuery(models.ClientId(r.ClientId))
}
