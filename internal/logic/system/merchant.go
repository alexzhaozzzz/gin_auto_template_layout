package system

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewMerchant() *Merchant {
	return &Merchant{}
}

type Merchant struct {
}

func (s Merchant) GetListByPage(c *ginx.Context) {
	req := &dto.Pagination{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Merchant GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	list, err := d.ListByPage(req)
	if err != nil {
		logrus.Errorf("Merchant GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount()
	if err != nil {
		logrus.Errorf("Merchant GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	req.TotalNum = int(total)

	resp := map[string]interface{}{"list": list, "page": req}
	c.RenderSuccess(resp)
	return
}

func (s Merchant) GetInfo(c *ginx.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Errorf("Menu GetInfo Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("Menu GetInfo Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("Merchant GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"info": info}
	c.RenderSuccess(resp)
	return
}

func (s Merchant) Add(c *ginx.Context) {
	req := &dto.MerchantList{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Merchant Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Merchant) Edit(c *ginx.Context) {
	req := &dto.MerchantList{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Merchant Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Merchant) Delete(c *ginx.Context) {
	req := &dto.MerchantList{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{
		Id: int32(req.Id),
	}
	d := system.NewMerchantData()
	err := d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
