// Package dto

package dto

type CurrencyRecordReq struct {
	Day      string `json:"day" form:"day"`             // 流水id
	TradeId  string `json:"trade_id" form:"trade_id"`   // 流水id
	PlayerId int64  `json:"player_id" form:"player_id"` // 玩家id
	Pagination
}
