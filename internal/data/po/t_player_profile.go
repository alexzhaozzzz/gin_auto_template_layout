// Package po

package po

import "time"

type PlayerProfile struct {
	NId                   int64     `gorm:"column:n_id;primaryKey;autoIncrement" json:"n_id"`
	NPlayerid             uint64    `gorm:"column:n_playerid" json:"n_playerid"`                             // 玩家id
	NType                 int32     `gorm:"column:n_type" json:"n_type"`                                     // 用户类型: 0=新手, 1=免费, 2=付费小r, 3=中r,4=大r，5=超r，33=黑名单用户
	NTag                  string    `gorm:"column:n_tag" json:"n_tag"`                                       // 用户标签列表, 用,分隔。
	NCreatetime           int64     `gorm:"column:n_createtime" json:"n_createtime"`                         // 创建时间 单位秒
	NScore                int64     `gorm:"column:n_score" json:"n_score"`                                   // 信任评分
	NUpdatetime           time.Time `gorm:"column:n_updatetime" json:"n_updatetime"`                         // 更新日期
	NRechargeCount        int32     `gorm:"column:n_recharge_count" json:"n_recharge_count"`                 // 充值次数
	NRechargeMax          int64     `gorm:"column:n_recharge_max" json:"n_recharge_max"`                     // 单次最大充值额
	NChannel              int32     `gorm:"column:n_channel" json:"n_channel"`                               // 渠道号
	NPrevPay              int64     `gorm:"column:n_prev_pay" json:"n_prev_pay"`                             // 上次充值金额
	NKind                 int32     `gorm:"column:n_kind" json:"n_kind"`                                     // 用户风险种类: 0=a普通, 1=b可疑，2=c安全
	NTotalChip            int64     `gorm:"column:n_total_chip" json:"n_total_chip"`                         // 总打码量
	NChipProgress         int64     `gorm:"column:n_chip_progress" json:"n_chip_progress"`                   // 打码进度
	NChipRequirement      int64     `gorm:"column:n_chip_requirement" json:"n_chip_requirement"`             // 打码需求
	NVipLevel             int32     `gorm:"column:n_vip_level" json:"n_vip_level"`                           // vip等级
	NOther                string    `gorm:"column:n_other" json:"n_other"`                                   // 其他信息记录
	NBonusChipProgress    int64     `gorm:"column:n_bonus_chip_progress" json:"n_bonus_chip_progress"`       // bonus打码进度
	NBonusChipRequirement int64     `gorm:"column:n_bonus_chip_requirement" json:"n_bonus_chip_requirement"` // bonus打码需求
	NMchId                int64     `gorm:"column:n_mch_id" json:"n_mch_id"`                                 // 商户id
}

func (p *PlayerProfile) TableName() string {
	return "t_player_profile"
}
