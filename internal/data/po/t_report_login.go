// Package po

package po

type ReportLogin struct {
	Id            int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Appid         int32  `gorm:"column:appid" json:"appid"`                  // 产品id
	Channelid     int32  `gorm:"column:channelid" json:"channel_id"`         // 渠道id
	Playerid      int64  `gorm:"column:playerid" json:"player_id"`           // 玩家id
	Nickname      string `gorm:"column:nickname" json:"nickname"`            // 昵称
	Type          int32  `gorm:"column:type" json:"type"`                    // 类型:0上线,1下线
	Ip            string `gorm:"column:ip" json:"ip"`                        // 登录ip
	Logindevice   string `gorm:"column:logindevice" json:"login_device"`     // 登陆设备
	Devicecode    string `gorm:"column:devicecode" json:"device_code"`       // 设备imei(唯一识别码)
	Clientversion string `gorm:"column:clientversion" json:"client_version"` // 客户端版本号
	Osinfomation  string `gorm:"column:osinfomation" json:"os_infomation"`   // 客户端系统信息
	Begintime     string `gorm:"column:begintime" json:"begin_time"`         // 登录时间
	Duration      int64  `gorm:"column:duration" json:"duration"`            // 在线时间:秒
	Endtime       string `gorm:"column:endtime" json:"end_time"`             // 离线时间
	Location      string `gorm:"column:location" json:"location"`            // 地区
}

func (r *ReportLogin) TableName() string {
	return "t_report_login"
}
