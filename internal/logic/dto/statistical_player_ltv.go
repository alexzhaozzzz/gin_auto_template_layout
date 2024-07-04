// Package dto

package dto

type StatisticalPlayerLtvReq struct {
	StartTime int64 `json:"start_time" form:"start_time"`
	EndTime   int64 `json:"end_time" form:"end_time"`
	Pagination
}

type ContentInfoInt struct {
	Day1  int `json:"day_1"`
	Day2  int `json:"day_2"`
	Day3  int `json:"day_3"`
	Day4  int `json:"day_4"`
	Day5  int `json:"day_5"`
	Day6  int `json:"day_6"`
	Day7  int `json:"day_7"`
	Day8  int `json:"day_8"`
	Day9  int `json:"day_9"`
	Day10 int `json:"day_10"`
	Day11 int `json:"day_11"`
	Day12 int `json:"day_12"`
	Day13 int `json:"day_13"`
	Day14 int `json:"day_14"`
	Day15 int `json:"day_15"`
	Day16 int `json:"day_16"`
	Day17 int `json:"day_17"`
	Day18 int `json:"day_18"`
	Day19 int `json:"day_19"`
	Day20 int `json:"day_20"`
	Day21 int `json:"day_21"`
	Day22 int `json:"day_22"`
	Day23 int `json:"day_23"`
	Day24 int `json:"day_24"`
	Day25 int `json:"day_25"`
	Day26 int `json:"day_26"`
	Day27 int `json:"day_27"`
	Day28 int `json:"day_28"`
	Day29 int `json:"day_29"`
	Day30 int `json:"day_30"`
	Day31 int `json:"day_31"`
	Day32 int `json:"day_32"`
	Day33 int `json:"day_33"`
	Day34 int `json:"day_34"`
	Day35 int `json:"day_35"`
	Day36 int `json:"day_36"`
	Day37 int `json:"day_37"`
	Day38 int `json:"day_38"`
	Day39 int `json:"day_39"`
	Day40 int `json:"day_40"`
	Day41 int `json:"day_41"`
	Day42 int `json:"day_42"`
	Day43 int `json:"day_43"`
	Day44 int `json:"day_44"`
	Day45 int `json:"day_45"`
	Day46 int `json:"day_46"`
	Day47 int `json:"day_47"`
	Day48 int `json:"day_48"`
	Day49 int `json:"day_49"`
	Day50 int `json:"day_50"`
	Day51 int `json:"day_51"`
	Day52 int `json:"day_52"`
	Day53 int `json:"day_53"`
	Day54 int `json:"day_54"`
	Day55 int `json:"day_55"`
	Day56 int `json:"day_56"`
	Day57 int `json:"day_57"`
	Day58 int `json:"day_58"`
	Day59 int `json:"day_59"`
	Day60 int `json:"day_60"`
}

type ContentInfoFloat64 struct {
	Day1  float64 `json:"day_1"`
	Day2  float64 `json:"day_2"`
	Day3  float64 `json:"day_3"`
	Day4  float64 `json:"day_4"`
	Day5  float64 `json:"day_5"`
	Day6  float64 `json:"day_6"`
	Day7  float64 `json:"day_7"`
	Day8  float64 `json:"day_8"`
	Day9  float64 `json:"day_9"`
	Day10 float64 `json:"day_10"`
	Day11 float64 `json:"day_11"`
	Day12 float64 `json:"day_12"`
	Day13 float64 `json:"day_13"`
	Day14 float64 `json:"day_14"`
	Day15 float64 `json:"day_15"`
	Day16 float64 `json:"day_16"`
	Day17 float64 `json:"day_17"`
	Day18 float64 `json:"day_18"`
	Day19 float64 `json:"day_19"`
	Day20 float64 `json:"day_20"`
	Day21 float64 `json:"day_21"`
	Day22 float64 `json:"day_22"`
	Day23 float64 `json:"day_23"`
	Day24 float64 `json:"day_24"`
	Day25 float64 `json:"day_25"`
	Day26 float64 `json:"day_26"`
	Day27 float64 `json:"day_27"`
	Day28 float64 `json:"day_28"`
	Day29 float64 `json:"day_29"`
	Day30 float64 `json:"day_30"`
	Day31 float64 `json:"day_31"`
	Day32 float64 `json:"day_32"`
	Day33 float64 `json:"day_33"`
	Day34 float64 `json:"day_34"`
	Day35 float64 `json:"day_35"`
	Day36 float64 `json:"day_36"`
	Day37 float64 `json:"day_37"`
	Day38 float64 `json:"day_38"`
	Day39 float64 `json:"day_39"`
	Day40 float64 `json:"day_40"`
	Day41 float64 `json:"day_41"`
	Day42 float64 `json:"day_42"`
	Day43 float64 `json:"day_43"`
	Day44 float64 `json:"day_44"`
	Day45 float64 `json:"day_45"`
	Day46 float64 `json:"day_46"`
	Day47 float64 `json:"day_47"`
	Day48 float64 `json:"day_48"`
	Day49 float64 `json:"day_49"`
	Day50 float64 `json:"day_50"`
	Day51 float64 `json:"day_51"`
	Day52 float64 `json:"day_52"`
	Day53 float64 `json:"day_53"`
	Day54 float64 `json:"day_54"`
	Day55 float64 `json:"day_55"`
	Day56 float64 `json:"day_56"`
	Day57 float64 `json:"day_57"`
	Day58 float64 `json:"day_58"`
	Day59 float64 `json:"day_59"`
	Day60 float64 `json:"day_60"`
}

type StatisticalPlayerLtvResp struct {
	Date     int32              `json:"date"`     // 日期总额
	Type     int32              `json:"type"`     // 枚举类型
	Register int32              `json:"register"` // 新增注册
	Recharge int64              `json:"recharge"` // 累计充值
	Withdraw int64              `json:"withdraw"` // 累计提现
	Content  ContentInfoFloat64 `json:"content"`  // 日期内容
}
