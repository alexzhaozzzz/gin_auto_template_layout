package logic

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
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

func (s Merchant) GetList(c *ginx.Context) {
	d := data.NewMerchantData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Merchant GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
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

	d := data.NewMerchantData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("Merchant GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"info": info}
	c.RenderSuccess(data)
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

	d := data.NewMerchantData()
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

	d := data.NewMerchantData()
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

	dReq := po.MerchantList{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Merchant Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewMerchantData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
