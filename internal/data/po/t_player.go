// Package po

package po

import "time"

type Player struct {
	Id            int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Accountid     int64     `gorm:"column:accountid" json:"accountid"` // 账户id
	Playerid      int64     `gorm:"column:playerid" json:"playerid"`   // 玩家id
	Showuid       int64     `gorm:"column:showuid" json:"showuid"`     // 显示id
	Type          int32     `gorm:"column:type" json:"type"`           // 玩家类型1.普通玩家，2.机器人，3.qa\n2.\n3.
	Channelid     int32     `gorm:"column:channelid" json:"channelid"` // 渠道id
	Productid     int32     `gorm:"column:productid" json:"productid"` // 产品id
	Inviteid      uint64    `gorm:"column:inviteid" json:"inviteid"`   // 邀请人玩家playerid
	Nickname      string    `gorm:"column:nickname" json:"nickname"`
	Gender        int32     `gorm:"column:gender" json:"gender"`               // 性别
	Avatar        string    `gorm:"column:avatar" json:"avatar"`               // 头像地址
	Hdavatar      string    `gorm:"column:hdavatar" json:"hdavatar"`           // 高清头像地址
	Provinceid    int32     `gorm:"column:provinceid" json:"provinceid"`       // 省id
	Cityid        int32     `gorm:"column:cityid" json:"cityid"`               // 市id
	Areaid        int32     `gorm:"column:areaid" json:"areaid"`               // 地区id
	Name          string    `gorm:"column:name" json:"name"`                   // 真实姓名
	Phone         string    `gorm:"column:phone" json:"phone"`                 // 手机号码
	Phonetype     int32     `gorm:"column:phonetype" json:"phonetype"`         // 手机类型:1.正常手机号，2.黑号
	Phonelocation int32     `gorm:"column:phonelocation" json:"phonelocation"` // 手机号归属地
	Idcard        string    `gorm:"column:idcard" json:"idcard"`               // 身份证
	Wxopenid      string    `gorm:"column:wxopenid" json:"wxopenid"`           // 游客账号id, 微信opendid或googleid，第三方账号id
	Wxunionid     string    `gorm:"column:wxunionid" json:"wxunionid"`         // 微信unionid
	Wxname        string    `gorm:"column:wxname" json:"wxname"`
	Iswhitelist   int32     `gorm:"column:iswhitelist" json:"iswhitelist"`   // 是否qa，默认否
	Zipcode       int32     `gorm:"column:zipcode" json:"zipcode"`           // 邮编
	Shippingaddr  string    `gorm:"column:shippingaddr" json:"shippingaddr"` // 收获地址
	Status        int32     `gorm:"column:status" json:"status"`             // 1可登录，2冻结，默认为1
	Remark        string    `gorm:"column:remark" json:"remark"`             // 备注
	Version       string    `gorm:"column:version" json:"version"`           // 客户端版本号
	Device        string    `gorm:"column:device" json:"device"`             // 客户端设备名称
	Imei          string    `gorm:"column:imei" json:"imei"`                 // 客户端设备唯一识别码
	Createtime    time.Time `gorm:"column:createtime" json:"createtime"`     // 创建时间，通常也是注册时间
	Createby      string    `gorm:"column:createby" json:"createby"`         // 创建人
	Updatetime    time.Time `gorm:"column:updatetime" json:"updatetime"`     // 更新时间
	Updateby      string    `gorm:"column:updateby" json:"updateby"`         // 更新人
	Promoterid    string    `gorm:"column:promoterid" json:"promoterid"`     // 推广员id
	Fromid        string    `gorm:"column:fromid" json:"fromid"`             // 来源id
	Zoneid        string    `gorm:"column:zoneid" json:"zoneid"`             // 场景id
	Topversion    string    `gorm:"column:topversion" json:"topversion"`     // 客户端最高版本号
	Regip         string    `gorm:"column:regip" json:"regip"`               // 注册ip
	Topinvite     int64     `gorm:"column:topinvite" json:"topinvite"`       // 最高上级玩家id
	MoneyPass     string    `gorm:"column:money_pass" json:"money_pass"`     // 玩家资金密码
	VipLevel      int32     `gorm:"column:vip_level" json:"vip_level"`       // vip等级
}

func (p *Player) TableName() string {
	return "t_player"
}
