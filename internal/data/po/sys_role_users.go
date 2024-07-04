// Package po

package po

type SysRoleUsers struct {
	Id         int32 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleId     int64 `gorm:"column:role_id" json:"role_id"`
	UserId     int64 `gorm:"column:user_id" json:"user_id"`
	CreateTime int32 `gorm:"column:create_time" json:"create_time"`
	UpdateTime int32 `gorm:"column:update_time" json:"update_time"`
}

func (s *SysRoleUsers) TableName() string {
	return "sys_role_users"
}
