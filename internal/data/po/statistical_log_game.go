// Package po

package po

import "time"

type StatisticalLogGame struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChannelId int32     `gorm:"column:channel_id" json:"channel_id"` // 渠道
	GameId    string    `gorm:"column:game_id" json:"game_id"`       // 游戏id
	Date      int32     `gorm:"column:date" json:"date"`             // 日期
	Platform  string    `gorm:"column:platform" json:"platform"`     // 平台
	JoinNum   int64     `gorm:"column:join_num" json:"join_num"`     // 参与人数
	Bet       int64     `gorm:"column:bet" json:"bet"`               // 下注
	Reward    int64     `gorm:"column:reward" json:"reward"`         // 返奖
	Win       int64     `gorm:"column:win" json:"win"`               // 输赢
	Chip      int64     `gorm:"column:chip" json:"chip"`             // 打码量
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (s *StatisticalLogGame) TableName() string {
	return "statistical_log_game"
}
