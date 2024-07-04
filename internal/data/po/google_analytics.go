// Package po

package po

import "time"

type GoogleAnalytics struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChannelId int32     `gorm:"column:channel_id" json:"channel_id"` // 渠道id
	Date      int32     `gorm:"column:date" json:"date"`             // 日期
	Pv        int32     `gorm:"column:pv" json:"pv"`                 // pv
	Uv        int32     `gorm:"column:uv" json:"uv"`                 // uv
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (g *GoogleAnalytics) TableName() string {
	return "google_analytics"
}
