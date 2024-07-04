// Package po

package po

import "time"

type PayFirstRecord struct {
	NId         int64     `gorm:"column:n_id;primaryKey;autoIncrement" json:"n_id"`
	NUid        int64     `gorm:"column:n_uid" json:"n_uid"`                 // 玩家id
	NChannel    int32     `gorm:"column:n_channel" json:"n_channel"`         // 渠道id
	NCoin       int64     `gorm:"column:n_coin" json:"n_coin"`               // 首次充值前的金币
	NAmount     string    `gorm:"column:n_amount" json:"n_amount"`           // 首次充值额
	NOrder      string    `gorm:"column:n_order" json:"n_order"`             // 首次充值订单id
	NCreateTime time.Time `gorm:"column:n_create_time" json:"n_create_time"` // 创建日期
}

func (p *PayFirstRecord) TableName() string {
	return "t_pay_first_record"
}
