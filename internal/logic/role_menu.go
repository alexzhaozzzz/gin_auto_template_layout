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

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{}
}

type RoleMenu struct {
}

func (s RoleMenu) GetList(c *ginx.Context) {
	d := conn.NewRoleMenuData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("RoleMenu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s RoleMenu) Add(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleMenu Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleMenu Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("RoleMenu Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RoleMenu) Edit(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleMenu Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleMenu Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("RoleMenu Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RoleMenu) Delete(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleMenu Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleMenu{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleMenu Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("RoleMenu Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
