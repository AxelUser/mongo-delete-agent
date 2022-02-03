package handler

import "github.com/AxelUser/mongo-delete-agent/src/models"

type job struct {
	id    int64
	query models.DataQuery
}
