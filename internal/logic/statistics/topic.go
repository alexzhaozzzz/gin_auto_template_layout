package statistics

import (
	"github.com/sirupsen/logrus"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/statistics"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewTopic() *Topic {
	return &Topic{}
}

type Topic struct {
}

func (s Topic) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Topic GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.StatisticalLogGame{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Topic GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Topic GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := statistics.NewTopicData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Topic GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("Topic GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": list, "page": pageResp}
	c.RenderSuccess(resp)
	return
}
