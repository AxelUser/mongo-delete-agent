package requests

type JobReq struct {
	JobId int64 `header:"jobId" binding:"required"`
}
