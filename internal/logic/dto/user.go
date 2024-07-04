// Package dto

package dto

type UserChangeReq struct {
	User     User      `json:"user"`
	RoleUser RoleUsers `json:"role_user"`
}

type User struct {
	Id         int64  `json:"id"`
	Guid       string `json:"guid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	NickName   string `json:"nick_name"`
	MerchantId int64  `json:"merchant_id"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type RoleUsers struct {
	RoleId int64 `json:"role_id"`
}
