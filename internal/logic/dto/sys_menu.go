// Package dto

package dto

type SysMenuReq struct {
	Id             int64   `json:"id"`
	ParentId       int64   `json:"parent_id"`       // 上级菜单
	Title          string  `json:"title"`           // 标题
	Icon           string  `json:"icon"`            // 图标
	Uri            string  `json:"uri"`             // 路径
	Show           int32   `json:"show"`            // 是否展示:1是,0否
	Sort           int32   `json:"sort"`            // 排序
	PermissionsIds []int64 `json:"permissions_ids"` // 权限表对应的ids json数组
}
