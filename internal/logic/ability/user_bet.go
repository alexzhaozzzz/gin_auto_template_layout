package ability

import (
	"github.com/sirupsen/logrus"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewUserBet() *UserBet {
	return &UserBet{}
}

type UserBet struct {
}

// GetListByPage ...
// @Summary 获取用户每日充提差
// @Tags todos
// @Produce  json
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param player_id query int true "用户ID"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.BetStatistics,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/userbet [get]
func (s UserBet) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("UserBet GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.UserBetReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("UserBet GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("UserBet GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewBetStatisticsData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("UserBet GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("UserBet GetListByPage ListCount Err: %s", err.Error())
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
