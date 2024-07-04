// Package po

package po

import "time"

type RechargeRecord struct {
	Id             uint64    `gorm:"column:n_id;primaryKey;autoIncrement" json:"id"`
	Orderid        string    `gorm:"column:n_orderid" json:"order_id"`                // 订单id
	Playerid       uint64    `gorm:"column:n_playerid" json:"player_id"`              // 玩家id
	Amount         string    `gorm:"column:n_amount" json:"amount"`                   // 充值金额，小数点2位
	State          int32     `gorm:"column:n_state" json:"state"`                     // 订单状态: 0=未付款, 1=已付款,2=已发货，3=补发成功， -1=充值失败，-2=发货失败
	Des            string    `gorm:"column:n_des" json:"des"`                         // 订单失败描述
	Createtime     time.Time `gorm:"column:n_createtime" json:"create_time"`          // 订单创建时间
	Paytime        time.Time `gorm:"column:n_paytime" json:"pay_time"`                // 订单付款时间
	Finishtime     time.Time `gorm:"column:n_finishtime" json:"finish_time"`          // 订单完成时间（发货成功或失败）
	Coin           int64     `gorm:"column:n_coin" json:"coin"`                       // 得到金币数
	Channel        int32     `gorm:"column:n_channel" json:"channel"`                 // 渠道号
	ExOrderid      string    `gorm:"column:n_ex_orderid" json:"ex_order_id"`          // 外部订单号
	Bankinfo       string    `gorm:"column:n_bankinfo" json:"bank_info"`              // 银行卡信息: 姓名和卡号, json
	ExResp         string    `gorm:"column:n_ex_resp" json:"ex_resp"`                 // 外部支付平台回复内容
	From           string    `gorm:"column:n_from" json:"from"`                       // 支付渠道
	Phone          string    `gorm:"column:n_phone" json:"phone"`                     // 手机号
	Name           string    `gorm:"column:n_name" json:"name"`                       // 姓名
	Partnerid      string    `gorm:"column:n_partnerid" json:"partner_id"`            // 商户id
	Platform       string    `gorm:"column:n_platform" json:"platform"`               // 充值平台
	GoodsId        int32     `gorm:"column:n_goods_id" json:"goods_id"`               // 充值商品id
	Pic            string    `gorm:"column:n_pic" json:"pic"`                         // 转账截图列表(json格式)
	Inviteid       int64     `gorm:"column:n_inviteid" json:"invite_id"`              // 邀请人玩家playerid
	Topinvite      int64     `gorm:"column:n_topinvite" json:"top_invite"`            // 最高上级玩家id
	Src            int32     `gorm:"column:n_src" json:"src"`                         // 来源id： 1=商城按钮， 2=+号，3=充值优惠按钮， 4=登陆奖励入口， 5=月卡vip， 6=匹配金币不足。
	InputType      int32     `gorm:"column:n_input_type" json:"input_type"`           // 1=选择，2=手动输入
	ClientType     int32     `gorm:"column:n_client_type" json:"client_type"`         // 客户端类型: 1=apk，2=html5
	ClientVer      string    `gorm:"column:n_client_ver" json:"client_ver"`           // 客户端版本号
	Entertime      time.Time `gorm:"column:n_entertime" json:"enter_time"`            // 进入商城页面时间
	BeforeCoin     int64     `gorm:"column:n_before_coin" json:"before_coin"`         // 充值前的金币数
	OpenSrc        int32     `gorm:"column:n_open_src" json:"open_src"`               // 打开来源id
	VpayCreatetime string    `gorm:"column:n_vpay_createtime" json:"pay_create_time"` // 回调createtime
	Red            int32     `gorm:"column:n_red" json:"red"`                         // 红点状态
}

func (r *RechargeRecord) TableName() string {
	return "t_recharge_record"
}
