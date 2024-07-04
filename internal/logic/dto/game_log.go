// Package dto

package dto

type GameLogReq struct {
	Day      string `json:"day" form:"day"`
	PlayerId int64  `json:"player_id" form:"player_id"` // 玩家id
	GameId   int32  `json:"game_id" form:"game_id"`     // 游戏id
	Bet      int64  `json:"bet" form:"bet"`             // 下注
	Reward   int64  `json:"reward" form:"reward"`       // 返奖
	Win      int64  `json:"win" form:"win"`
	LogType  int32  `json:"log_type" json:"log_type"`
	Pagination
}
