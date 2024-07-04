// Package dto

package dto

type StatisticalLogGameReq struct {
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
	Pagination
}

type TopicResp struct {
	Id        int64   `json:"id"`
	ChannelId int32   `json:"channel_id"` // 渠道
	GameId    string  `json:"game_id"`    // 游戏id
	Date      int32   `json:"date"`       // 日期
	Platform  string  `json:"platform"`   // 平台
	JoinNum   int64   `json:"join_num"`   // 参与人数
	Bet       int64   `json:"bet"`        // 下注
	Reward    int64   `json:"reward"`     // 返奖
	Win       int64   `json:"win"`        // 输赢
	Chip      int64   `json:"chip"`       // 打码量
	Rtp       float64 `json:"rtp"`
}
