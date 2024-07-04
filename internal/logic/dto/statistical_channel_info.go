// Package dto

package dto

type StatisticalChannelInfoReq struct {
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
	Pagination
}
