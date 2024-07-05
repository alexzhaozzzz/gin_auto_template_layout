// Package dto

package dto

type SysMenuReq struct {
	Id             int64   `json:"id" form:"id"`
	ParentId       int64   `json:"parent_id" form:"parent_id"`             // 上级菜单
	Title          string  `json:"title" form:"title"`                     // 标题
	Icon           string  `json:"icon" form:"icon"`                       // 图标
	Uri            string  `json:"uri" form:"uri"`                         // 路径
	Show           int32   `json:"show" form:"show"`                       // 是否展示:1是,0否
	Sort           int32   `json:"sort" form:"sort"`                       // 排序
	PermissionsIds []int64 `json:"permissions_ids" form:"permissions_ids"` // 权限表对应的ids json数组
}
