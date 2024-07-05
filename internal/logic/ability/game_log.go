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

func NewGameLog() *GameLog {
	return &GameLog{}
}

type GameLog struct {
}

// GetListByPage ...
// @Summary 用户游戏日志
// @Tags Ability
// @Produce  json
// @Param day query string true "日期（例:20240501）"
// @Param player_id query int false "玩家id"
// @Param game_id query int false "游戏id"
// @Param bet query int false "下注"
// @Param reward query int false "返奖"
// @Param win query int false "赢取"
// @Param log_type query int false "日志分类 0=下注返奖日志 1=下注日志 2=返奖日志"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.GameLog,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/gamelog [get]
func (s GameLog) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("GameLogReq GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.GameLogReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("GameLogReq GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("GameLogReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if req.Day == "" {
		now := time.Now()
		req.Day = fmt.Sprintf("%d%d%d", now.YearDay(), now.Month(), now.Day())
	}

	d := ability.NewGameLogData()
	list, err := d.ListByPage(req.Day, cIds, req)
	if err != nil {
		logrus.Errorf("GameLogReq GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(req.Day, cIds, req)
	if err != nil {
		logrus.Errorf("GameLogReq GetListByPage ListCount Err: %s", err.Error())
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
