package ability

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"
	"time"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewGoldChanges() *GoldChanges {
	return &GoldChanges{}
}

type GoldChanges struct {
}

// GetListByPage ...
// @Summary 每日金币变化
// @Tags Ability
// @Produce  json
// @Param day query string true "日期（例:20240501）"
// @Param trade_id query string false "流水id"
// @Param player_id query int false "玩家id"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.CurrencyRecord,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/goldchange [get]
func (s GoldChanges) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("GoldChanges GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.CurrencyRecordReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("GoldChanges GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("GoldChanges GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if req.Day == "" {
		now := time.Now()
		req.Day = fmt.Sprintf("%d%d%d", now.YearDay(), now.Month(), now.Day())
	}

	d := ability.NewGoldChangesData()
	list, err := d.ListByPage(req.Day, cIds, req)
	if err != nil {
		logrus.Errorf("GoldChanges GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(req.Day, cIds, req)
	if err != nil {
		logrus.Errorf("GoldChanges GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := ginx.ListResponses{
		List: list,
		Page: pageResp,
	}
	c.RenderSuccess(resp)
	return
}
