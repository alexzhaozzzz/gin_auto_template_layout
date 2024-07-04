// Package po

package po

type BetStatistics struct {
	ChannelId int32 `gorm:"column:channel_id" json:"channel_id"` // 渠道id
	PlayerId  int64 `gorm:"column:player_id" json:"player_id"`   // 玩家id
	GameId    int32 `gorm:"column:game_id" json:"game_id"`       // 游戏id
	Day       int64 `gorm:"column:day" json:"day"`               // 日期(当日时间戳)
	Times     int64 `gorm:"column:times" json:"times"`           // 当日总次数
	Bet       int64 `gorm:"column:bet" json:"bet"`               // 当日总下注
	Reward    int64 `gorm:"column:reward" json:"reward"`         // 当日总返奖
	Fee       int64 `gorm:"column:fee" json:"fee"`               // 当日总抽水
	Win       int64 `gorm:"column:win" json:"win"`               // 当日总输赢
	Chip      int64 `gorm:"column:chip" json:"chip"`             // 当日打码量
}

func (b *BetStatistics) TableName() string {
	return "bet_statistics"
}
