// Package po

package po

type MerchantList struct {
	Id         int32  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Guid       string `gorm:"column:guid" json:"guid"`
	Name       string `gorm:"column:name" json:"name"` //  商户名称
	Host       string `gorm:"column:host" json:"host"` // 站点域名
	ManageHost string `gorm:"column:manage_host" json:"manage_host"`
	Status     int32  `gorm:"column:status" json:"status"`
	CreateTime int32  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int32  `gorm:"column:update_time" json:"update_time"`
}

func (m *MerchantList) TableName() string {
	return "merchant_list"
}
