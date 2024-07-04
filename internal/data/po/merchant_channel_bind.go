// Package po

package po

type MerchantChannelBind struct {
	Id         int32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MerchantId int32 `gorm:"column:merchant_id" json:"merchant_id"` // 商户 id
	ChannelId  int32 `gorm:"column:channel_id" json:"channel_id"`   // 渠道 id
	Status     int32 `gorm:"column:status" json:"status"`
	CreateTime int32 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int32 `gorm:"column:update_time" json:"update_time"`
}

func (m *MerchantChannelBind) TableName() string {
	return "merchant_channel_bind"
}
