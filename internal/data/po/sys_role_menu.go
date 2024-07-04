// Package po

package po

type SysRoleMenu struct {
	Id         int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleId     int64 `gorm:"column:role_id" json:"role_id"`
	MenuId     int64 `gorm:"column:menu_id" json:"menu_id"`
	CreateTime int32 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int32 `gorm:"column:update_time" json:"update_time"`
}

func (s *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
