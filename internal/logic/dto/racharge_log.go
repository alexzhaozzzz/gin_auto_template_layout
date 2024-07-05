// Package dto

package dto

type RechargeLogReq struct {
	StartTime       int64  `json:"start_time" form:"start_time"`
	EndTime         int64  `json:"end_time" form:"end_time"`
	FinishStartTime int64  `json:"finish_start_time" form:"finish_start_time"`
	FinishEndTime   int64  `json:"finish_end_time" form:"finish_end_time"`
	PlayerId        int64  `json:"player_id" form:"player_id"` // 玩家id
	OrderId         string `json:"order_id" form:"order_id"`
	Amount          string `json:"amount" form:"amount"`
	State           int32  `json:"state" form:"state"`
	Platform        string `json:"platform" form:"platform"`
	Pagination
}

type RechargeLogResp struct {
	Id             uint64 `json:"id"`
	Orderid        string `json:"order_id"`        // 订单id
	Playerid       uint64 `json:"player_id"`       // 玩家id
	Amount         string `json:"amount"`          // 充值金额，小数点2位
	State          int32  `json:"state"`           // 订单状态: 0=未付款, 1=已付款,2=已发货，3=补发成功， -1=充值失败，-2=发货失败
	Des            string `json:"des"`             // 订单失败描述
	Createtime     string `json:"create_time"`     // 订单创建时间
	Paytime        string `json:"pay_time"`        // 订单付款时间
	Finishtime     string `json:"finish_time"`     // 订单完成时间（发货成功或失败）
	Coin           int64  `json:"coin"`            // 得到金币数
	Channel        int32  `json:"channel"`         // 渠道号
	ExOrderid      string `json:"ex_order_id"`     // 外部订单号
	Bankinfo       string `json:"bank_info"`       // 银行卡信息: 姓名和卡号, json
	ExResp         string `json:"ex_resp"`         // 外部支付平台回复内容
	From           string `json:"from"`            // 支付渠道
	Phone          string `json:"phone"`           // 手机号
	Name           string `json:"name"`            // 姓名
	Partnerid      string `json:"partner_id"`      // 商户id
	Platform       string `json:"platform"`        // 充值平台
	GoodsId        int32  `json:"goods_id"`        // 充值商品id
	Pic            string `json:"pic"`             // 转账截图列表(json格式)
	Inviteid       int64  `json:"invite_id"`       // 邀请人玩家playerid
	Topinvite      int64  `json:"top_invite"`      // 最高上级玩家id
	Src            int32  `json:"src"`             // 来源id： 1=商城按钮， 2=+号，3=充值优惠按钮， 4=登陆奖励入口， 5=月卡vip， 6=匹配金币不足。
	InputType      int32  `json:"input_type"`      // 1=选择，2=手动输入
	ClientType     int32  `json:"client_type"`     // 客户端类型: 1=apk，2=html5
	ClientVer      string `json:"client_ver"`      // 客户端版本号
	Entertime      string `json:"enter_time"`      // 进入商城页面时间
	BeforeCoin     int64  `json:"before_coin"`     // 充值前的金币数
	OpenSrc        int32  `json:"open_src"`        // 打开来源id
	VpayCreatetime string `json:"pay_create_time"` // 回调createtime
	Red            int32  `json:"red"`             // 红点状态
}
