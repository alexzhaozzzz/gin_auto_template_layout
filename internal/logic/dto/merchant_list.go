// Package dto

package dto

type MerchantListReq struct {
	Id         int64  `json:"id"`
	Guid       string `json:"guid"`
	Name       string `json:"name"` //  商户名称
	Host       string `json:"host"` // 站点域名
	ManageHost string `json:"manage_host"`
	Status     int32  `json:"status"`
	CreateTime int32  `json:"create_time"`
	UpdateTime int32  `json:"update_time"`
}
