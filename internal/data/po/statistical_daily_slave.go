// Package po

package po

type StatisticalDailySlave struct {
	SDId                       uint32  `gorm:"column:s_d_id;primaryKey" json:"s_d_id"`                                      // 1:1 statistical_daily (id)
	RegisterContent            string  `gorm:"column:register_content" json:"register_content"`                             // 注册玩家id
	LoginContent               string  `gorm:"column:login_content" json:"login_content"`                                   // 登录玩家id
	GoldContent                string  `gorm:"column:gold_content" json:"gold_content"`                                     // 活跃玩家id
	SpinActive                 int32   `gorm:"column:spin_active" json:"spin_active"`                                       // 新增活跃
	Arpu                       int32   `gorm:"column:arpu" json:"arpu"`                                                     // arpu活跃帐号的充值金额
	NewArpu                    int32   `gorm:"column:new_arpu" json:"new_arpu"`                                             // 新用户的当天充值金额
	OldArpu                    int32   `gorm:"column:old_arpu" json:"old_arpu"`                                             // 老用户(非当天注册)当天充值金额
	OldArpuRecharge            int32   `gorm:"column:old_arpu_recharge" json:"old_arpu_recharge"`                           // 老用户(非当天首充)当天充值金额
	ActiveWithdraw             int32   `gorm:"column:active_withdraw" json:"active_withdraw"`                               // 活跃提现
	NewWithdraw                int32   `gorm:"column:new_withdraw" json:"new_withdraw"`                                     // 新用户提现
	OldWithdraw                int32   `gorm:"column:old_withdraw" json:"old_withdraw"`                                     // 老用户提现
	Coins                      int64   `gorm:"column:coins" json:"coins"`                                                   // 非绑定金币
	Ingots                     int64   `gorm:"column:ingots" json:"ingots"`                                                 // 绑定金币
	IngotsPay                  int64   `gorm:"column:ingots_pay" json:"ingots_pay"`                                         // 付费用户绑金
	NBonusSum                  int64   `gorm:"column:n_bonus_sum" json:"n_bonus_sum"`                                       // 奖金
	NewPayLimit                int64   `gorm:"column:new_pay_limit" json:"new_pay_limit"`                                   // 新增充值额度
	RechargeContent            string  `gorm:"column:recharge_content" json:"recharge_content"`                             // 付费充值的玩家id
	WithdrawContent            string  `gorm:"column:withdraw_content" json:"withdraw_content"`                             // 提现玩家id
	FirstRechargeContent       string  `gorm:"column:first_recharge_content" json:"first_recharge_content"`                 // 首充玩家id
	Bankruptcy                 int64   `gorm:"column:bankruptcy" json:"bankruptcy"`                                         // 破产玩家数量
	NewPayContent              string  `gorm:"column:new_pay_content" json:"new_pay_content"`                               // 新注册付费玩家
	ReNewPayContent            string  `gorm:"column:re_new_pay_content" json:"re_new_pay_content"`                         // 复购新注册付费玩家
	NewDeviceLoginNum          int32   `gorm:"column:new_device_login_num" json:"new_device_login_num"`                     // 新增设备登录
	OldDeviceLoginNum          int32   `gorm:"column:old_device_login_num" json:"old_device_login_num"`                     // 老设备登录
	RechargeSuccessRate        float64 `gorm:"column:recharge_success_rate" json:"recharge_success_rate"`                   // 支付成功率
	NewRechargeSuccessRate     float64 `gorm:"column:new_recharge_success_rate" json:"new_recharge_success_rate"`           // 新用户增支付成功率
	OldRechargeSuccessRate     float64 `gorm:"column:old_recharge_success_rate" json:"old_recharge_success_rate"`           // 老用户付费成功率
	NewRechargeSuccessRateCurr float64 `gorm:"column:new_recharge_success_rate_curr" json:"new_recharge_success_rate_curr"` // 新用户增支付成功率(新增用户支付成功订单/ 新增用户所有支付订单)
	OldRechargeSuccessRateCurr float64 `gorm:"column:old_recharge_success_rate_curr" json:"old_recharge_success_rate_curr"` // 老用户付费成功率(老用户支付成功订单/老用户所有支付订单)
	InstallSrcContent          string  `gorm:"column:install_src_content" json:"install_src_content"`                       // 安装来源
	NotSdkContent              string  `gorm:"column:not_sdk_content" json:"not_sdk_content"`                               // 非sdk真金模式用户
	ReOldPayContent            string  `gorm:"column:re_old_pay_content" json:"re_old_pay_content"`                         // 复购老用户
}

func (s *StatisticalDailySlave) TableName() string {
	return "statistical_daily_slave"
}
