// Package po

package po

import "time"

type PlayerDevice struct {
	NId            int64     `gorm:"column:n_id;primaryKey;autoIncrement" json:"n_id"`
	NPlayerId      int64     `gorm:"column:n_player_id" json:"n_player_id"`             // 玩家id
	NChannel       int32     `gorm:"column:n_channel" json:"n_channel"`                 // 渠道id
	NRegIp         string    `gorm:"column:n_reg_ip" json:"n_reg_ip"`                   // 注册ip
	NClientVer     string    `gorm:"column:n_client_ver" json:"n_client_ver"`           // 客户端版本号
	NOs            string    `gorm:"column:n_os" json:"n_os"`                           // 系统信息
	NDevCode       string    `gorm:"column:n_dev_code" json:"n_dev_code"`               // 设备唯一id
	NLocation      string    `gorm:"column:n_location" json:"n_location"`               // 玩家所在地区
	NCoordinate    string    `gorm:"column:n_coordinate" json:"n_coordinate"`           // 坐标
	NLang          string    `gorm:"column:n_lang" json:"n_lang"`                       // 手机系统语言
	NPhoneName     string    `gorm:"column:n_phone_name" json:"n_phone_name"`           // 手机型号
	NSdkData       string    `gorm:"column:n_sdk_data" json:"n_sdk_data"`               // leo sdk data
	NSp            string    `gorm:"column:n_sp" json:"n_sp"`                           // 手机运营商
	NIsReal        int32     `gorm:"column:n_is_real" json:"n_is_real"`                 // 是否是真金用户
	NAdid          string    `gorm:"column:n_adid" json:"n_adid"`                       // ad id
	NRegTime       time.Time `gorm:"column:n_reg_time" json:"n_reg_time"`               // 注册时间
	NLastLoginTime int64     `gorm:"column:n_last_login_time" json:"n_last_login_time"` // 最后登录时间，unix时间戳
}

func (p *PlayerDevice) TableName() string {
	return "t_player_device"
}
