package ability

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewLvt() *Lvt {
	return &Lvt{}
}

type Lvt struct {
}

func (s Lvt) compute(c, t int, r int32) float64 {
	var resp float64

	if c+t > 0 && r > 0 {
		resp = float64((c + t) / int(r))
	}
	return resp
}

// GetListByPage ...
// @Summary 盈亏LTV
// @Tags Statistics
// @Produce  json
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.StatisticalPlayerLtvResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /statistics/lvt [get]
func (s Lvt) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Lvt GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.StatisticalPlayerLtvReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Lvt GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Lvt GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewLvtData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Lvt GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]dto.StatisticalPlayerLtvResp, 0)
	if len(list) > 0 {
		for _, v := range list {
			var dayContent dto.ContentInfoInt
			var dayDrawContent dto.ContentInfoInt

			if v.DayContent != "" {
				// 使用json.Unmarshal进行解码
				err := json.Unmarshal([]byte(v.DayContent), &dayContent)
				if err != nil {
				}
			}

			if v.DayContentWithdraw != "" {
				// 使用json.Unmarshal进行解码
				err := json.Unmarshal([]byte(v.DayContentWithdraw), &dayDrawContent)
				if err != nil {
				}
			}

			contentInfo := dto.ContentInfoFloat64{
				Day1:  s.compute(dayContent.Day1, dayDrawContent.Day1, v.Register),
				Day2:  s.compute(dayContent.Day2, dayDrawContent.Day2, v.Register),
				Day3:  s.compute(dayContent.Day3, dayDrawContent.Day3, v.Register),
				Day4:  s.compute(dayContent.Day4, dayDrawContent.Day4, v.Register),
				Day5:  s.compute(dayContent.Day5, dayDrawContent.Day5, v.Register),
				Day6:  s.compute(dayContent.Day6, dayDrawContent.Day6, v.Register),
				Day7:  s.compute(dayContent.Day7, dayDrawContent.Day7, v.Register),
				Day8:  s.compute(dayContent.Day8, dayDrawContent.Day8, v.Register),
				Day9:  s.compute(dayContent.Day9, dayDrawContent.Day9, v.Register),
				Day10: s.compute(dayContent.Day10, dayDrawContent.Day10, v.Register),
				Day11: s.compute(dayContent.Day11, dayDrawContent.Day11, v.Register),
				Day12: s.compute(dayContent.Day12, dayDrawContent.Day12, v.Register),
				Day13: s.compute(dayContent.Day13, dayDrawContent.Day13, v.Register),
				Day14: s.compute(dayContent.Day14, dayDrawContent.Day14, v.Register),
				Day15: s.compute(dayContent.Day15, dayDrawContent.Day15, v.Register),
				Day16: s.compute(dayContent.Day16, dayDrawContent.Day16, v.Register),
				Day17: s.compute(dayContent.Day17, dayDrawContent.Day17, v.Register),
				Day18: s.compute(dayContent.Day18, dayDrawContent.Day18, v.Register),
				Day19: s.compute(dayContent.Day19, dayDrawContent.Day19, v.Register),
				Day20: s.compute(dayContent.Day20, dayDrawContent.Day20, v.Register),
				Day21: s.compute(dayContent.Day21, dayDrawContent.Day21, v.Register),
				Day22: s.compute(dayContent.Day22, dayDrawContent.Day22, v.Register),
				Day23: s.compute(dayContent.Day23, dayDrawContent.Day23, v.Register),
				Day24: s.compute(dayContent.Day24, dayDrawContent.Day24, v.Register),
				Day25: s.compute(dayContent.Day25, dayDrawContent.Day25, v.Register),
				Day26: s.compute(dayContent.Day26, dayDrawContent.Day26, v.Register),
				Day27: s.compute(dayContent.Day27, dayDrawContent.Day27, v.Register),
				Day28: s.compute(dayContent.Day28, dayDrawContent.Day28, v.Register),
				Day29: s.compute(dayContent.Day29, dayDrawContent.Day29, v.Register),
				Day30: s.compute(dayContent.Day30, dayDrawContent.Day30, v.Register),
				Day31: s.compute(dayContent.Day31, dayDrawContent.Day31, v.Register),
				Day32: s.compute(dayContent.Day32, dayDrawContent.Day32, v.Register),
				Day33: s.compute(dayContent.Day33, dayDrawContent.Day33, v.Register),
				Day34: s.compute(dayContent.Day34, dayDrawContent.Day34, v.Register),
				Day35: s.compute(dayContent.Day35, dayDrawContent.Day35, v.Register),
				Day36: s.compute(dayContent.Day36, dayDrawContent.Day36, v.Register),
				Day37: s.compute(dayContent.Day37, dayDrawContent.Day37, v.Register),
				Day38: s.compute(dayContent.Day38, dayDrawContent.Day38, v.Register),
				Day39: s.compute(dayContent.Day39, dayDrawContent.Day39, v.Register),
				Day40: s.compute(dayContent.Day40, dayDrawContent.Day40, v.Register),
				Day41: s.compute(dayContent.Day41, dayDrawContent.Day41, v.Register),
				Day42: s.compute(dayContent.Day42, dayDrawContent.Day42, v.Register),
				Day43: s.compute(dayContent.Day43, dayDrawContent.Day43, v.Register),
				Day44: s.compute(dayContent.Day44, dayDrawContent.Day44, v.Register),
				Day45: s.compute(dayContent.Day45, dayDrawContent.Day45, v.Register),
				Day46: s.compute(dayContent.Day46, dayDrawContent.Day46, v.Register),
				Day47: s.compute(dayContent.Day47, dayDrawContent.Day47, v.Register),
				Day48: s.compute(dayContent.Day48, dayDrawContent.Day48, v.Register),
				Day49: s.compute(dayContent.Day49, dayDrawContent.Day49, v.Register),
				Day50: s.compute(dayContent.Day50, dayDrawContent.Day50, v.Register),
				Day51: s.compute(dayContent.Day51, dayDrawContent.Day51, v.Register),
				Day52: s.compute(dayContent.Day52, dayDrawContent.Day52, v.Register),
				Day53: s.compute(dayContent.Day53, dayDrawContent.Day53, v.Register),
				Day54: s.compute(dayContent.Day54, dayDrawContent.Day54, v.Register),
				Day55: s.compute(dayContent.Day55, dayDrawContent.Day55, v.Register),
				Day56: s.compute(dayContent.Day56, dayDrawContent.Day56, v.Register),
				Day57: s.compute(dayContent.Day57, dayDrawContent.Day57, v.Register),
				Day58: s.compute(dayContent.Day58, dayDrawContent.Day58, v.Register),
				Day59: s.compute(dayContent.Day59, dayDrawContent.Day59, v.Register),
				Day60: s.compute(dayContent.Day60, dayDrawContent.Day60, v.Register),
			}

			respList = append(respList, dto.StatisticalPlayerLtvResp{
				Date:     v.Date,
				Type:     v.Type,
				Recharge: v.Recharge,
				Withdraw: v.Withdraw,
				Register: v.Register,
				Content:  contentInfo,
			})
		}
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("Lvt GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := ginx.ListResponses{
		List: respList,
		Page: pageResp,
	}
	c.RenderSuccess(resp)
	return
}
