package dto

type ThirdOrderCheckReq struct {
	OrderId string `json:"order_id"`
}

type ToThirdOrderCheckReq struct {
	Server string                `json:"Server"`
	Cmd    string                `json:"Cmd"`
	Data   ToThirdOrderCheckData `json:"Data"`
}

type ToThirdOrderCheckData struct {
	OrderNo string `json:"orderNo"`
	Pass    string `json:"pass"`
	Desc    string `json:"desc"`
	User    string `json:"user"`
}

type ToThirdOrderCheckResp struct {
	OrderNo string `json:"orderNo"`  // 查询订单id
	Status  int    `json:"status"`   // 订单状态: 0 待支付 1 已支付 3 支付失败 4 支付异常, -1 查询失败
	PayId   string `json:"pay_id"`   // 充值通道ID
	PayName string `json:"pay_name"` // 充值通道名
	Desc    string `json:"desc"`     // 描述
}
