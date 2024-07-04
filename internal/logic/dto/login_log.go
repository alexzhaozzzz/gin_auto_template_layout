// Package dto

package dto

type LoginLogReq struct {
	//Day      string `json:"day" form:"day"`             // 流水id
	PlayerId int64 `json:"player_id" form:"player_id"` // 玩家id
	Pagination
}

type LoginLogResp struct {
	Appid            int32  `json:"appid"`              // 产品id
	Channelid        int32  `json:"channel_id"`         // 渠道id
	Playerid         int64  `json:"player_id"`          // 玩家id
	Nickname         string `json:"nickname"`           // 昵称
	Type             int32  `json:"type"`               // 类型:0上线,1下线
	Ip               string `json:"ip"`                 // 登录ip
	Logindevice      string `json:"login_device"`       // 登陆设备
	Devicecode       string `json:"device_code"`        // 设备imei(唯一识别码)
	Clientversion    string `json:"client_version"`     // 客户端版本号
	Osinfomation     string `json:"os_infomation"`      // 客户端系统信息
	Begintime        string `json:"begin_time"`         // 登录时间
	Duration         int64  `json:"duration"`           // 在线时间:秒
	Endtime          string `json:"end_time"`           // 离线时间
	Location         string `json:"location"`           // 地区
	FirstPayNum      string `json:"first_pay_num"`      // 首充金额
	FirstPayTime     int64  `json:"first_pay_time"`     // 首充时间
	FirstPayoutsNum  string `json:"first_payouts_num"`  // 首提金额
	FirstPayoutsTime int64  `json:"first_payouts_time"` // 首提时间
}
