package logic

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	conn "gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewMenu() *Menu {
	return &Menu{}
}

type Menu struct {
}

func (s Menu) GetList(c *ginx.Context) {
	d := conn.NewMenuData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Menu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Menu) Add(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Menu Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewMenuData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Menu Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Menu) Edit(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Menu Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewMenuData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Menu Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Menu) Delete(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Menu Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewMenuData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Menu Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
