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

func NewRetain() *Retain {
	return &Retain{}
}

type Retain struct {
}

// GetListByPage ...
// @Summary 留存
// @Tags Statistics
// @Produce  json
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param r_type query int false "类型: 1登录,2注册,3充值,4提现"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.Retain,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /statistics/retain [get]
func (s Retain) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Retain GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.RetainReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Retain GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Retain GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewRetainData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Retain GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("Retain GetListByPage ListCount Err: %s", err.Error())
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
