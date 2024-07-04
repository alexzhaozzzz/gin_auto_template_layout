// Package po

package po

type SysRoles struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Guid       string `gorm:"column:guid" json:"guid"`
	MerchantId int64  `gorm:"column:merchant_id" json:"merchant_id"` // 商户id
	Name       string `gorm:"column:name" json:"name"`               // 角色名称
	Code       string `gorm:"column:code" json:"code"`               // 角色标识
	CreateTime int32  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int32  `gorm:"column:update_time" json:"update_time"`
}

func (s *SysRoles) TableName() string {
	return "sys_roles"
}
