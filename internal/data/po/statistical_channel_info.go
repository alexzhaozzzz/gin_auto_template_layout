// Package po

package po

import "time"

type StatisticalChannelInfo struct {
	Id                  int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ChannelId           int32     `gorm:"column:channel_id" json:"channel_id"`                       // 渠道
	Date                int32     `gorm:"column:date" json:"date"`                                   // 日期
	NewDev              int64     `gorm:"column:new_dev" json:"new_dev"`                             // 新增设备
	Reg                 int64     `gorm:"column:reg" json:"reg"`                                     // 新增注册
	RegBind             int64     `gorm:"column:reg_bind" json:"reg_bind"`                           // 新增绑定
	TotalBind           int64     `gorm:"column:total_bind" json:"total_bind"`                       // 总绑定
	RegRecharge         int64     `gorm:"column:reg_recharge" json:"reg_recharge"`                   // 新充用户
	RegRechargeTimes    int64     `gorm:"column:reg_recharge_times" json:"reg_recharge_times"`       // 新充次数
	RegRechargeAmount   int64     `gorm:"column:reg_recharge_amount" json:"reg_recharge_amount"`     // 新充金额
	RegWithdrawAmount   int64     `gorm:"column:reg_withdraw_amount" json:"reg_withdraw_amount"`     // 新提金额
	FirstRecharge       int64     `gorm:"column:first_recharge" json:"first_recharge"`               // 首充用户
	TotalRecharge       int64     `gorm:"column:total_recharge" json:"total_recharge"`               // 总充用户
	TotalRechargeTimes  int64     `gorm:"column:total_recharge_times" json:"total_recharge_times"`   // 总充次数
	TotalRechargeAmount int64     `gorm:"column:total_recharge_amount" json:"total_recharge_amount"` // 总充值金额
	TotalWithdrawAmount int64     `gorm:"column:total_withdraw_amount" json:"total_withdraw_amount"` // 总提现金额
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`                       // 创建时间
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`                       // 更新时间
	RegRechargeBind     int32     `gorm:"column:reg_recharge_bind" json:"reg_recharge_bind"`         // 新绑充用户
}

func (s *StatisticalChannelInfo) TableName() string {
	return "statistical_channel_info"
}
