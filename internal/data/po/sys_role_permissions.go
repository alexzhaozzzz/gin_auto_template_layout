// Package po

package po

type SysRolePermissions struct {
	Id           int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleId       int64 `gorm:"column:role_id" json:"role_id"`
	PermissionId int64 `gorm:"column:permission_id" json:"permission_id"`
	CreateTime   int32 `gorm:"column:create_time" json:"create_time"`
	UpdateTime   int32 `gorm:"column:update_time" json:"update_time"`
}

func (s *SysRolePermissions) TableName() string {
	return "sys_role_permissions"
}
