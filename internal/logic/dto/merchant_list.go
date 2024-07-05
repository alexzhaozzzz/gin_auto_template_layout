// Package dto

package dto

type MerchantListReq struct {
	Id         int64  `json:"id" form:"id"`
	Guid       string `json:"guid" form:"guid"`
	Name       string `json:"name" form:"name"` //  商户名称
	Host       string `json:"host" form:"host"` // 站点域名
	ManageHost string `json:"manage_host" form:"manage_host"`
	Status     int32  `json:"status" form:"status"`
	CreateTime int32  `json:"create_time" form:"create_time"`
	UpdateTime int32  `json:"update_time" form:"update_time"`
}
