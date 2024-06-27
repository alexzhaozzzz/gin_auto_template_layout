package billing

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/billing"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewPayouts() *Payouts {
	return &Payouts{}
}

type Payouts struct {
}

func (s Payouts) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Payouts GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.Payouts{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Payouts GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Payouts GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := billing.NewPayoutsData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Payouts GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("Payouts GetListByPage ListCount Err: %s", err.Error())
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

func (s Payouts) GetAuditListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Payouts GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.Payouts{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Payouts GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Payouts GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := billing.NewPayoutsData()
	list, err := d.ListByStatePage(cIds, 0, req)
	if err != nil {
		logrus.Errorf("Payouts GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListStateCount(cIds, 0)
	if err != nil {
		logrus.Errorf("Payouts GetListByPage ListCount Err: %s", err.Error())
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
