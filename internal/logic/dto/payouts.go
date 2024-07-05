// Package dto

package dto

type PayoutsReq struct {
	PlayerId       int64  `json:"player_id" form:"player_id"` // 玩家id
	OrderId        string `json:"order_id" form:"order_id"`   //订单id
	State          int32  `json:"state" form:"state"`         // 订单状态: 0=审核中, 1=审核成功,2=提现成功,-1=审核失败,-2=提现失败,-3=提现提交失败
	OrderStartTime int64  `json:"order_start_time" form:"order_start_time"`
	OrderEndTime   int64  `json:"order_end_time" form:"order_end_time"`
	LoanStartTime  int64  `json:"loan_start_time" form:"loan_start_time"`
	LoanEndTime    int64  `json:"loan_end_time" form:"loan_end_time"`
	Pagination
}

type PayoutsAuditReq struct {
	PlayerId    int64  `json:"player_id" form:"player_id"` // 玩家id
	OrderId     string `json:"order_id" form:"order_id"`
	AmountStart int64  `json:"amount_start" form:"amount_start"`
	AmountEnd   int64  `json:"amount_end" form:"amount_end"`
	Pagination
}

type PayoutsAuditResp struct {
	Id            uint64 `json:"id"`
	Orderid       string `json:"order_id"`        // 订单id
	Playerid      uint64 `json:"player_id"`       // 玩家id
	Amount        string `json:"amount"`          // 提现金额，小数点2位
	State         int32  `json:"state"`           // 订单状态: 0=审核中, 1=审核成功,2=提现成功,-1=审核失败,-2=提现失败,-3=提现提交失败
	Des           string `json:"des"`             // 订单失败描述
	Createtime    string `json:"create_time"`     // 订单创建时间
	Paytime       string `json:"pay_time"`        // 订单放款时间
	Checktime     string `json:"check_time"`      // 订单审核时间
	Coin          int64  `json:"coin"`            // 扣除金币数
	Channel       int32  `json:"channel"`         // 用户渠道号
	ExOrderid     string `json:"ex_order_id"`     // 外部订单号
	Bankinfo      string `json:"bank_info"`       // 银行卡信息: 姓名和卡号, json
	ExResp        string `json:"ex_resp"`         // 外部支付平台回复内容
	BackCoin      int32  `json:"back_coin"`       // 提现失败，返回金币是否成功: 1=成功, -1=失败,0=未返回。
	Operator      string `json:"operator"`        // 操作人
	Fee           string `json:"fee"`             // 提现费用
	From          string `json:"from"`            // 来源
	Partnerid     string `json:"partner_id"`      // 商户id
	Platform      string `json:"platform"`        // 充值平台
	After         int64  `json:"after"`           // 提现之后的金额
	Kind          int32  `json:"kind"`            // 提现银行卡类型: upi, mps
	ShowErr       string `json:"show_err"`        // 客户端显示错误
	CardId        string `json:"card_id"`         // 银行卡信息
	PayCreatetime string `json:"pay_create_time"` // 回调createtime
	Retry         int32  `json:"retry"`           // 重试次数
	RetryState    int32  `json:"retry_state"`     // 重试提交状态: 0=未检测, 1=自动提交,2=手动审核
	QueryState    int32  `json:"query_state"`     // 重试提交状态: 0=未检测, 1=自动提交,2=手动审核
	QueryResp     string `json:"query_resp"`      // 查询回复
	LastOrderId   string `json:"last_order_id"`   // 最后的订单id
	Red           int32  `json:"red"`             // 红点状态
	WdFlag        int32  `json:"wd_flag"`         // 提现标识 0普通 1代理
	UserFlag      int32  `json:"user_flag"`       // 0普通 1优质
	UserKind      int32  `json:"user_kind"`       // 用户风险种类: 0=a普通, 1=b可疑，2=c安全
}

type UserBaseReq struct {
	PlayerId int64 `json:"player_id" form:"player_id"` // 玩家id
}

type UserBaseResp struct {
	PlayerId   int64  `json:"player_id"` // 玩家id
	NickName   string `json:"nick_name"`
	Card       string `json:"card"`
	CardName   string `json:"card_name"`
	Phone      string `json:"phone"`
	DeviceCode string `json:"device_code"`
	Tag        string `json:"tag"`
	UType      int32  `json:"u_type"`
}

type UserRechargeListReq struct {
	PlayerId       int64 `json:"player_id" form:"player_id"` // 玩家id
	OrderStartTime int64 `json:"order_start_time" form:"order_start_time"`
	OrderEndTime   int64 `json:"order_end_time" form:"order_end_time"`
	Pagination
}
