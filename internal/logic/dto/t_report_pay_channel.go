// Package dto

package dto

type ReportPayChannelReq struct {
	Pid int32 `json:"pid"` // 通道id
	Pagination
}
