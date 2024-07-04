// Package po

package po

type SysMenu struct {
	Id             int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ParentId       int64  `gorm:"column:parent_id" json:"parent_id"`             // 上级菜单
	Title          string `gorm:"column:title" json:"title"`                     // 标题
	Icon           string `gorm:"column:icon" json:"icon"`                       // 图标
	Uri            string `gorm:"column:uri" json:"uri"`                         // 路径
	Show           int32  `gorm:"column:show" json:"show"`                       // 是否展示:1是,0否
	Sort           int32  `gorm:"column:sort" json:"sort"`                       // 排序
	PermissionsIds string `gorm:"column:permissions_ids" json:"permissions_ids"` // 权限表对应的ids json数组
	CreateTime     int32  `gorm:"column:create_time" json:"create_time"`
	UpdateTime     int32  `gorm:"column:update_time" json:"update_time"`
}

func (s *SysMenu) TableName() string {
	return "sys_menu"
}
