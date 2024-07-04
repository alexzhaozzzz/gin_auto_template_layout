// Package po

package po

type GameLog struct {
	Id            int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 自增id
	RecordId      string `gorm:"column:record_id" json:"record_id"`            // 记录id
	PlayerId      int64  `gorm:"column:player_id" json:"player_id"`            // 玩家id
	ChannelId     int32  `gorm:"column:channel_id" json:"channel_id"`          // 渠道id
	PlatformId    int32  `gorm:"column:platform_id" json:"platform_id"`        // 游戏平台id
	GameId        int32  `gorm:"column:game_id" json:"game_id"`                // 游戏id
	GameType      int32  `gorm:"column:game_type" json:"game_type"`            // 游戏类型 0 slot 1 百人场 2 牌局 3 崩溃类
	LevelId       int32  `gorm:"column:level_id" json:"level_id"`              // 场次id
	DeskId        int64  `gorm:"column:desk_id" json:"desk_id"`                // 桌子id
	DeskType      int32  `gorm:"column:desk_type" json:"desk_type"`            // 桌子类型
	RegTime       int64  `gorm:"column:reg_time" json:"reg_time"`              // 注册时间
	ChipRule      int32  `gorm:"column:chip_rule" json:"chip_rule"`            // 打码规则 0:不打码 1:slot 2:brc 3:crash 4:poker
	Chip          int64  `gorm:"column:chip" json:"chip"`                      // 此局打码量
	Balance       int64  `gorm:"column:balance" json:"balance"`                // 游戏后总余额
	Cash          int64  `gorm:"column:cash" json:"cash"`                      // 游戏后真金
	Deposit       int64  `gorm:"column:deposit" json:"deposit"`                // 游戏后绑金
	Bonus         int64  `gorm:"column:bonus" json:"bonus"`                    // 游戏后bonus
	BeforeBalance int64  `gorm:"column:before_balance" json:"before_balance"`  // 游戏前总余额
	BeforeCash    int64  `gorm:"column:before_cash" json:"before_cash"`        // 游戏前真金
	BeforeDeposit int64  `gorm:"column:before_deposit" json:"before_deposit"`  // 游戏前绑金
	BeforeBonus   int64  `gorm:"column:before_bonus" json:"before_bonus"`      // 游戏前bonus
	Recharge      int64  `gorm:"column:recharge" json:"recharge"`              // 总充值
	Withdrawal    int64  `gorm:"column:withdrawal" json:"withdrawal"`          // 总提现
	Bet           int64  `gorm:"column:bet" json:"bet"`                        // 下注
	AddCash       int64  `gorm:"column:add_cash" json:"add_cash"`              // 赢真金
	AddDeposit    int64  `gorm:"column:add_deposit" json:"add_deposit"`        // 赢绑金
	Reward        int64  `gorm:"column:reward" json:"reward"`                  // 返奖
	Win           int64  `gorm:"column:win" json:"win"`                        // 输赢值=返奖-下注-抽水
	Fee           int64  `gorm:"column:fee" json:"fee"`                        // 抽水
	IsFree        int32  `gorm:"column:is_free" json:"is_free"`                // 是否免费
	WinType       int32  `gorm:"column:win_type" json:"win_type"`              // 赢奖类型 0=未赢，1=小赢，2=big win，3=mega win
	PlanId        int32  `gorm:"column:plan_id" json:"plan_id"`                // 计划id
	ControlLog    string `gorm:"column:control_log" json:"control_log"`        // 游戏日志
	ResultLog     string `gorm:"column:result_log" json:"result_log"`          // 游戏结果
	ExtLog        string `gorm:"column:ext_log" json:"ext_log"`                // 游戏结果额外记录
	StartTime     int64  `gorm:"column:start_time" json:"start_time"`          // 游戏开始时间
	EndTime       int64  `gorm:"column:end_time" json:"end_time"`              // 游戏结束时间
	LinkGameLog   string `gorm:"column:link_game_log" json:"link_game_log"`    // 游戏结果
	LogType       int32  `gorm:"column:log_type" json:"log_type"`              // 日志分类 0=下注返奖日志 1=下注日志 2=返奖日志
}

func (g *GameLog) TableName() string {
	return "game_log"
}
