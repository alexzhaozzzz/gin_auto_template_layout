// Package dto

package dto

type StatisticalDailyReq struct {
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
	NewReg    int64 `json:"new_reg" form:"new_reg"`
	Recharge  int64 `json:"recharge" form:"recharge"`
	Pagination
}

type StatisticalDailyResp struct {
	Register                   uint32  `json:"register"`        // 日期区间内注册的账号数量
	Login                      uint32  `json:"login"`           // 日期区间内有过登录行为的账号数量
	Active                     int32   `json:"active"`          // 日期区间内参与spin人数
	Prepaid                    int32   `json:"prepaid"`         // 日期区间内充值总额
	PrepaidCount               int32   `json:"prepaid_count"`   // 充值笔数
	PrepaidNumber              int32   `json:"prepaid_number"`  // 日期区间内充值人数
	Withdraw                   int32   `json:"withdraw"`        // 日期区间内已提现成功金额
	WithdrawNew                int32   `json:"withdraw_new"`    // 当天首次提现的人数
	WithdrawNumber             int32   `json:"withdraw_number"` // 日期区间内已提现成功人数
	NewPay                     int32   `json:"new_pay"`         // 新增付费人数+1
	Win                        uint64  `json:"win"`
	Bet                        uint64  `json:"bet"`
	BetRealTime                int64   `json:"bet_real_time"`                  // 实时投注
	WinRealTime                int64   `json:"win_real_time"`                  // 实时赢奖情况
	Coins                      int64   `json:"coins"`                          // 可提金币
	Type                       int32   `json:"type"`                           // models
	TimeInt                    int32   `json:"time_int"`                       // 当天时间缀
	OldPay                     int32   `json:"old_pay"`                        // 付费老用户
	ReOldPay                   float64 `json:"re_old_pay"`                     // 老用户复购率
	FirstRecharge              int32   `json:"first_recharge"`                 // 当日首充人数
	FirstRechargeTotal         int64   `json:"first_recharge_total"`           // 当日首充用户充值总额
	ChannelId                  int32   `json:"channel_id"`                     // 渠道id default(0)
	RegBlackList               int32   `json:"reg_black_list"`                 // 当日注册关联黑名单数量
	ActiveAvgSpin              int32   `json:"active_avg_spin"`                // 新注册活跃平均spin次数
	ActiveAvgBet               int64   `json:"active_avg_bet"`                 // 新注册活跃平均下注
	TimeTag                    int32   `json:"time_tag"`                       // 小时标识
	Installs                   uint32  `json:"installs"`                       // 总安装量
	FbInstalls                 int32   `json:"fb_installs"`                    // facebook渠道安装量
	InstallsUpdateStatus       int32   `json:"installs_update_status"`         // 查询adjust安装信息状态 -1失败 0未更新 1更新成功
	RegisterContent            string  `json:"register_content"`               // 注册玩家id
	LoginContent               string  `json:"login_content"`                  // 登录玩家id
	GoldContent                string  `json:"gold_content"`                   // 活跃玩家id
	SpinActive                 int32   `json:"spin_active"`                    // 新增活跃
	Arpu                       int32   `json:"arpu"`                           // arpu活跃帐号的充值金额
	NewArpu                    int32   `json:"new_arpu"`                       // 新用户的当天充值金额
	OldArpu                    int32   `json:"old_arpu"`                       // 老用户(非当天注册)当天充值金额
	OldArpuRecharge            int32   `json:"old_arpu_recharge"`              // 老用户(非当天首充)当天充值金额
	ActiveWithdraw             int32   `json:"active_withdraw"`                // 活跃提现
	NewWithdraw                int32   `json:"new_withdraw"`                   // 新用户提现
	OldWithdraw                int32   `json:"old_withdraw"`                   // 老用户提现
	NoBindCoins                int64   `json:"no_bind_coins"`                  // 非绑定金币
	Ingots                     int64   `json:"ingots"`                         // 绑定金币
	IngotsPay                  int64   `json:"ingots_pay"`                     // 付费用户绑金
	NBonusSum                  int64   `json:"n_bonus_sum"`                    // 奖金
	NewPayLimit                int64   `json:"new_pay_limit"`                  // 新增充值额度
	RechargeContent            string  `json:"recharge_content"`               // 付费充值的玩家id
	WithdrawContent            string  `json:"withdraw_content"`               // 提现玩家id
	FirstRechargeContent       string  `json:"first_recharge_content"`         // 首充玩家id
	Bankruptcy                 int64   `json:"bankruptcy"`                     // 破产玩家数量
	NewPayContent              string  `json:"new_pay_content"`                // 新注册付费玩家
	ReNewPayContent            string  `json:"re_new_pay_content"`             // 复购新注册付费玩家
	NewDeviceLoginNum          int32   `json:"new_device_login_num"`           // 新增设备登录
	OldDeviceLoginNum          int32   `json:"old_device_login_num"`           // 老设备登录
	RechargeSuccessRate        float64 `json:"recharge_success_rate"`          // 支付成功率
	NewRechargeSuccessRate     float64 `json:"new_recharge_success_rate"`      // 新用户增支付成功率
	OldRechargeSuccessRate     float64 `json:"old_recharge_success_rate"`      // 老用户付费成功率
	NewRechargeSuccessRateCurr float64 `json:"new_recharge_success_rate_curr"` // 新用户增支付成功率(新增用户支付成功订单/ 新增用户所有支付订单)
	OldRechargeSuccessRateCurr float64 `json:"old_recharge_success_rate_curr"` // 老用户付费成功率(老用户支付成功订单/老用户所有支付订单)
	InstallSrcContent          string  `json:"install_src_content"`            // 安装来源
	NotSdkContent              string  `json:"not_sdk_content"`                // 非sdk真金模式用户
	ReOldPayContent            string  `json:"re_old_pay_content"`             // 复购老用户
	Pv                         int32   `json:"pv"`
	Uv                         int32   `json:"uv"`
}
