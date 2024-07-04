// Package po

package po

import "time"

type PlayerBank struct {
	NId           int64     `gorm:"column:n_id;primaryKey;autoIncrement" json:"n_id"`
	NPlayerid     uint64    `gorm:"column:n_playerid" json:"n_playerid"`         // 玩家id
	NName         string    `gorm:"column:n_name" json:"n_name"`                 // 姓名
	NBankcardid   string    `gorm:"column:n_bankcardid" json:"n_bankcardid"`     // 银行卡号或upi或imps账号
	NBankname     string    `gorm:"column:n_bankname" json:"n_bankname"`         // 银行名称
	NBankifsc     string    `gorm:"column:n_bankifsc" json:"n_bankifsc"`         // 银行ifsc
	NPhone        string    `gorm:"column:n_phone" json:"n_phone"`               // 手机号
	NCreatetime   time.Time `gorm:"column:n_createtime" json:"n_createtime"`     // 创建时间
	NIsselect     int32     `gorm:"column:n_isselect" json:"n_isselect"`         // 是否选中，默认银行卡： 1=默认,0=非默认
	NRazcontactid string    `gorm:"column:n_razcontactid" json:"n_razcontactid"` // raz联系人id
	NEmail        string    `gorm:"column:n_email" json:"n_email"`               // 邮箱
	NPanno        string    `gorm:"column:n_panno" json:"n_panno"`
	NAdno         string    `gorm:"column:n_adno" json:"n_adno"`
	NKind         int32     `gorm:"column:n_kind" json:"n_kind"`               // 银行卡类型: 0=银行卡, 1=upi,2=imps
	NUpdateTime   string    `gorm:"column:n_update_time" json:"n_update_time"` // 修改时间
}

func (p *PlayerBank) TableName() string {
	return "t_player_bank"
}
