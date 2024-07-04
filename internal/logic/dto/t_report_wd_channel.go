// Package dto

package dto

type ReportWdChannelReq struct {
	Pid int32 `json:"pid"` // 通道id
	Pagination
}
