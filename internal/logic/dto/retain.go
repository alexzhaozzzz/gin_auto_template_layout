// Package dto

package dto

type RetainReq struct {
	RType     int64 `json:"r_type" form:"r_type"`
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
	Pagination
}
