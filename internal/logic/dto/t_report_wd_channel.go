// Package dto

package dto

type ReportWdChannelReq struct {
	Pid int32 `json:"pid" form:"pid"` // 通道id
	Pagination
}
