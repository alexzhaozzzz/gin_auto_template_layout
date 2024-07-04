// Package po

package po

import "time"

type StatisticalDaily struct {
	Id                   uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Register             uint32    `gorm:"column:register" json:"register"`               // 日期区间内注册的账号数量
	Login                uint32    `gorm:"column:login" json:"login"`                     // 日期区间内有过登录行为的账号数量
	Active               int32     `gorm:"column:active" json:"active"`                   // 日期区间内参与spin人数
	Prepaid              int32     `gorm:"column:prepaid" json:"prepaid"`                 // 日期区间内充值总额
	PrepaidCount         int32     `gorm:"column:prepaid_count" json:"prepaid_count"`     // 充值笔数
	PrepaidNumber        int32     `gorm:"column:prepaid_number" json:"prepaid_number"`   // 日期区间内充值人数
	Withdraw             int32     `gorm:"column:withdraw" json:"withdraw"`               // 日期区间内已提现成功金额
	WithdrawNew          int32     `gorm:"column:withdraw_new" json:"withdraw_new"`       // 当天首次提现的人数
	WithdrawNumber       int32     `gorm:"column:withdraw_number" json:"withdraw_number"` // 日期区间内已提现成功人数
	NewPay               int32     `gorm:"column:new_pay" json:"new_pay"`                 // 新增付费人数+1
	Win                  uint64    `gorm:"column:win" json:"win"`
	Bet                  uint64    `gorm:"column:bet" json:"bet"`
	BetRealTime          int64     `gorm:"column:bet_real_time" json:"bet_real_time"`                   // 实时投注
	WinRealTime          int64     `gorm:"column:win_real_time" json:"win_real_time"`                   // 实时赢奖情况
	Coins                int64     `gorm:"column:coins" json:"coins"`                                   // 可提金币
	Type                 int32     `gorm:"column:type" json:"type"`                                     // models
	TimeInt              int32     `gorm:"column:time_int" json:"time_int"`                             // 当天时间缀
	OldPay               int32     `gorm:"column:old_pay" json:"old_pay"`                               // 付费老用户
	ReOldPay             float64   `gorm:"column:re_old_pay" json:"re_old_pay"`                         // 老用户复购率
	FirstRecharge        int32     `gorm:"column:first_recharge" json:"first_recharge"`                 // 当日首充人数
	FirstRechargeTotal   int64     `gorm:"column:first_recharge_total" json:"first_recharge_total"`     // 当日首充用户充值总额
	ChannelId            int32     `gorm:"column:channel_id" json:"channel_id"`                         // 渠道id default(0)
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`                         // 创建时间
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`                         // 更新时间
	RegBlackList         int32     `gorm:"column:reg_black_list" json:"reg_black_list"`                 // 当日注册关联黑名单数量
	ActiveAvgSpin        int32     `gorm:"column:active_avg_spin" json:"active_avg_spin"`               // 新注册活跃平均spin次数
	ActiveAvgBet         int64     `gorm:"column:active_avg_bet" json:"active_avg_bet"`                 // 新注册活跃平均下注
	TimeTag              int32     `gorm:"column:time_tag" json:"time_tag"`                             // 小时标识
	Installs             uint32    `gorm:"column:installs" json:"installs"`                             // 总安装量
	FbInstalls           int32     `gorm:"column:fb_installs" json:"fb_installs"`                       // facebook渠道安装量
	InstallsUpdateStatus int32     `gorm:"column:installs_update_status" json:"installs_update_status"` // 查询adjust安装信息状态 -1失败 0未更新 1更新成功
}

func (s *StatisticalDaily) TableName() string {
	return "statistical_daily"
}
