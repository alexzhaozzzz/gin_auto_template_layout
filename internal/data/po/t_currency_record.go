// Package po

package po

import "time"

type CurrencyRecord struct {
	Id            int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Tradeid       string    `gorm:"column:tradeid" json:"trade_id"`             // 流水id
	Playerid      int64     `gorm:"column:playerid" json:"player_id"`           // 玩家id
	Channel       int32     `gorm:"column:channel" json:"channel"`              // 渠道id
	Currencytype  int32     `gorm:"column:currencytype" json:"currency_type"`   // 货币类型: 1=金币, 2=元宝（钻石）， 3=房卡
	Amount        int64     `gorm:"column:amount" json:"amount"`                // 加减值
	Beforebalance int64     `gorm:"column:beforebalance" json:"before_balance"` // 操作前金币值
	Afterbalance  int64     `gorm:"column:afterbalance" json:"after_balance"`   // 操作后金币值
	Tradetime     time.Time `gorm:"column:tradetime" json:"trade_time"`         // 创建时间
	Status        int32     `gorm:"column:status" json:"status"`                // 操作结果： 1=成功，0=失败
	Remark        string    `gorm:"column:remark" json:"remark"`                // 备注
	Gameid        int64     `gorm:"column:gameid" json:"game_id"`               // 游戏id
	Level         int32     `gorm:"column:level" json:"level"`                  // 场次id
	Funcid        int32     `gorm:"column:funcid" json:"func_id"`               // 行为id或功能id
	Productid     int64     `gorm:"column:productid" json:"product_id"`         // 产品id。 红幺鸡=9999，小程序=10000
}

func (c *CurrencyRecord) TableName() string {
	return "t_currency_record"
}
