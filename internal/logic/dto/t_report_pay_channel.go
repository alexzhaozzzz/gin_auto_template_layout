// Package dto

package dto

type ReportPayChannelReq struct {
	Pid int32 `json:"pid" form:"pid"` // 通道id
	Pagination
}
