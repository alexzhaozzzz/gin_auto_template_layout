// Package dto

package dto

type SysPermissionsReq struct {
	Id         int64  `json:"id"`
	MenuId     int64  `json:"menu_id"` // 权限归属菜单
	Name       string `json:"name"`    // 权限名称
	Code       string `json:"code"`    // 标识
	Path       string `json:"path"`
	Command    string `json:"command"`
	CreateTime int32  `json:"create_time"`
	UpdateTime int32  `json:"update_time"`
}
