// Package po

package po

import "time"

type StatisticalRankPay struct {
	Id               int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChannelId        int64     `gorm:"column:channel_id" json:"channel_id"`               // 渠道id
	PlayerId         int64     `gorm:"column:player_id" json:"player_id"`                 // 玩家id
	TimeInt          int32     `gorm:"column:time_int" json:"time_int"`                   // 当天时间缀
	Pay              int64     `gorm:"column:pay" json:"pay"`                             // 当天累计充值
	Withdraw         int64     `gorm:"column:withdraw" json:"withdraw"`                   // 当日累计提现
	PayNumber        int32     `gorm:"column:pay_number" json:"pay_number"`               // 充值笔数
	Gold             int64     `gorm:"column:gold" json:"gold"`                           // 真金+绑金
	CurrencyRecharge int64     `gorm:"column:currency_recharge" json:"currency_recharge"` // 充值总额_当天
	CurrencyWithdraw int64     `gorm:"column:currency_withdraw" json:"currency_withdraw"` // 总提现额_当天
	NSrc             int32     `gorm:"column:n_src" json:"n_src"`                         // 充值来源（最后一笔充值
	NLastTheme       int32     `gorm:"column:n_last_theme" json:"n_last_theme"`           // spin主题（最后一笔充值前spin主题）
	Begintime        string    `gorm:"column:begintime" json:"begintime"`                 // 最后登录
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
}

func (s *StatisticalRankPay) TableName() string {
	return "statistical_rank_pay"
}
