// Package po

package po

import "time"

type WithdrawRecord struct {
	Id            uint64    `gorm:"column:n_id;primaryKey;autoIncrement" json:"id"`
	Orderid       string    `gorm:"column:n_orderid" json:"order_id"`                // 订单id
	Playerid      uint64    `gorm:"column:n_playerid" json:"player_id"`              // 玩家id
	Amount        string    `gorm:"column:n_amount" json:"amount"`                   // 提现金额，小数点2位
	State         int32     `gorm:"column:n_state" json:"state"`                     // 订单状态: 0=审核中, 1=审核成功,2=提现成功,-1=审核失败,-2=提现失败,-3=提现提交失败
	Des           string    `gorm:"column:n_des" json:"des"`                         // 订单失败描述
	Createtime    time.Time `gorm:"column:n_createtime" json:"create_time"`          // 订单创建时间
	Paytime       time.Time `gorm:"column:n_paytime" json:"pay_time"`                // 订单放款时间
	Checktime     time.Time `gorm:"column:n_checktime" json:"check_time"`            // 订单审核时间
	Coin          int64     `gorm:"column:n_coin" json:"coin"`                       // 扣除金币数
	Channel       int32     `gorm:"column:n_channel" json:"channel"`                 // 用户渠道号
	ExOrderid     string    `gorm:"column:n_ex_orderid" json:"ex_order_id"`          // 外部订单号
	Bankinfo      string    `gorm:"column:n_bankinfo" json:"bank_info"`              // 银行卡信息: 姓名和卡号, json
	ExResp        string    `gorm:"column:n_ex_resp" json:"ex_resp"`                 // 外部支付平台回复内容
	BackCoin      int32     `gorm:"column:n_back_coin" json:"back_coin"`             // 提现失败，返回金币是否成功: 1=成功, -1=失败,0=未返回。
	Operator      string    `gorm:"column:n_operator" json:"operator"`               // 操作人
	Fee           string    `gorm:"column:n_fee" json:"fee"`                         // 提现费用
	From          string    `gorm:"column:n_from" json:"from"`                       // 来源
	Partnerid     string    `gorm:"column:n_partnerid" json:"partner_id"`            // 商户id
	Platform      string    `gorm:"column:n_platform" json:"platform"`               // 充值平台
	After         int64     `gorm:"column:n_after" json:"after"`                     // 提现之后的金额
	Kind          int32     `gorm:"column:n_kind" json:"kind"`                       // 提现银行卡类型: upi, mps
	ShowErr       string    `gorm:"column:n_show_err" json:"show_err"`               // 客户端显示错误
	CardId        string    `gorm:"column:n_card_id" json:"card_id"`                 // 银行卡信息
	PayCreatetime string    `gorm:"column:n_vpay_createtime" json:"pay_create_time"` // 回调createtime
	Retry         int32     `gorm:"column:n_retry" json:"retry"`                     // 重试次数
	RetryState    int32     `gorm:"column:n_retry_state" json:"retry_state"`         // 重试提交状态: 0=未检测, 1=自动提交,2=手动审核
	QueryState    int32     `gorm:"column:n_query_state" json:"query_state"`         // 重试提交状态: 0=未检测, 1=自动提交,2=手动审核
	QueryResp     string    `gorm:"column:n_query_resp" json:"query_resp"`           // 查询回复
	LastOrderId   string    `gorm:"column:n_last_order_id" json:"last_order_id"`     // 最后的订单id
	Red           int32     `gorm:"column:n_red" json:"red"`                         // 红点状态
	WdFlag        int32     `gorm:"column:n_wd_flag" json:"wd_flag"`                 // 提现标识 0普通 1代理
	UserFlag      int32     `gorm:"column:n_user_flag" json:"user_flag"`             // 0普通 1优质
}

func (w *WithdrawRecord) TableName() string {
	return "t_withdraw_record"
}
