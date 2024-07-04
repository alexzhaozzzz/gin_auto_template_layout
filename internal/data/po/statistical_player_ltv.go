// Package po

package po

import "time"

type StatisticalPlayerLtv struct {
	Id                 int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChannelId          int32     `gorm:"column:channel_id" json:"channel_id"`                     // 渠道
	Date               int32     `gorm:"column:date" json:"date"`                                 // 日期总额
	Type               int32     `gorm:"column:type" json:"type"`                                 // 枚举类型
	Register           int32     `gorm:"column:register" json:"register"`                         // 新增注册
	Recharge           int64     `gorm:"column:recharge" json:"recharge"`                         // 累计充值
	Withdraw           int64     `gorm:"column:withdraw" json:"withdraw"`                         // 累计提现
	DayContent         string    `gorm:"column:day_content" json:"day_content"`                   // 日期内容(充值)
	DayContentWithdraw string    `gorm:"column:day_content_withdraw" json:"day_content_withdraw"` // 提现内容
	RegisterContent    string    `gorm:"column:register_content" json:"register_content"`         // 注册人员列表
	CostUsd            int32     `gorm:"column:cost_usd" json:"cost_usd"`                         // 冗余cost_data.cost_usd的数据,便于/roi的筛选查询
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`                     // 创建时间
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`                     // 更新时间
}

func (s *StatisticalPlayerLtv) TableName() string {
	return "statistical_player_ltv"
}
