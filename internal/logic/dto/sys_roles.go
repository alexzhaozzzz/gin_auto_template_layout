// Package dto
// Generated by sql2struct-0.0.5 at 2024-06-19 07:59:34
package dto

// SysRoles Generated by sql2struct-0.0.5 at 2024-06-19 07:59:34
type SysRoles struct {
    Id int64 `json:"id"`
    Guid string `json:"guid"`
    MerchantId int64 `json:"merchant_id"` // 商户id
    Name string `json:"name"` // 角色名称
    Code string `json:"code"` // 角色标识
    CreateTime int32 `json:"create_time"`
    UpdateTime int32 `json:"update_time"`
}
