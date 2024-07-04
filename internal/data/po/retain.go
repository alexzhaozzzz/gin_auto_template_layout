// Package po

package po

import "time"

type Retain struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // id
	ChannelId uint32    `gorm:"column:channel_id" json:"channel_id"`          // 日期
	Date      uint32    `gorm:"column:date" json:"date"`                      // 日期
	Type      uint32    `gorm:"column:type" json:"type"`                      // 类型: 1登录,2注册,3充值,4提现
	Num       uint32    `gorm:"column:num" json:"num"`                        // 人数
	R2        float64   `gorm:"column:r_2" json:"r_2"`                        // 留存
	R3        float64   `gorm:"column:r_3" json:"r_3"`                        // 留存
	R4        float64   `gorm:"column:r_4" json:"r_4"`                        // 留存
	R5        float64   `gorm:"column:r_5" json:"r_5"`                        // 留存
	R6        float64   `gorm:"column:r_6" json:"r_6"`                        // 留存
	R7        float64   `gorm:"column:r_7" json:"r_7"`                        // 留存
	R14       float64   `gorm:"column:r_14" json:"r_14"`                      // 留存
	R30       float64   `gorm:"column:r_30" json:"r_30"`                      // 留存
	R45       float64   `gorm:"column:r_45" json:"r_45"`                      // 留存
	R60       float64   `gorm:"column:r_60" json:"r_60"`                      // 留存
	R2Num     uint32    `gorm:"column:r_2_num" json:"r_2_num"`                // 留存数
	R3Num     uint32    `gorm:"column:r_3_num" json:"r_3_num"`                // 留存数
	R4Num     uint32    `gorm:"column:r_4_num" json:"r_4_num"`                // 留存数
	R5Num     uint32    `gorm:"column:r_5_num" json:"r_5_num"`                // 留存数
	R6Num     uint32    `gorm:"column:r_6_num" json:"r_6_num"`                // 留存数
	R7Num     uint32    `gorm:"column:r_7_num" json:"r_7_num"`                // 留存数
	R14Num    uint32    `gorm:"column:r_14_num" json:"r_14_num"`              // 留存数
	R30Num    uint32    `gorm:"column:r_30_num" json:"r_30_num"`              // 留存数
	R45Num    uint32    `gorm:"column:r_45_num" json:"r_45_num"`              // 留存数
	R60Num    uint32    `gorm:"column:r_60_num" json:"r_60_num"`              // 留存数
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`          // 更新时间
}

func (r *Retain) TableName() string {
	return "retain"
}
