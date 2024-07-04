// Package dto

package dto

type RechargeLogReq struct {
	StartTime       int64  `json:"start_time" form:"start_time"`
	EndTime         int64  `json:"end_time" form:"end_time"`
	FinishStartTime int64  `json:"finish_start_time" form:"finish_start_time"`
	FinishEndTime   int64  `json:"finish_end_time" form:"finish_end_time"`
	PlayerId        int64  `json:"player_id" form:"player_id"` // 玩家id
	OrderId         string `json:"order_id" form:"order_id"`
	Amount          string `json:"amount" form:"amount"`
	State           int32  `json:"state" form:"state"`
	Platform        string `json:"platform" form:"platform"`
	Pagination
}
