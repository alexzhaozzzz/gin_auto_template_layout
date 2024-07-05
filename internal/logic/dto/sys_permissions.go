// Package dto

package dto

type SysPermissionsReq struct {
	Id         int64  `json:"id" form:"id"`
	MenuId     int64  `json:"menu_id" form:"menu_id"` // 权限归属菜单
	Name       string `json:"name" form:"name"`       // 权限名称
	Code       string `json:"code" form:"code"`       // 标识
	Path       string `json:"path" form:"path"`
	Command    string `json:"command" form:"command"`
	CreateTime int32  `json:"create_time" form:"create_time"`
	UpdateTime int32  `json:"update_time" form:"update_time"`
}
