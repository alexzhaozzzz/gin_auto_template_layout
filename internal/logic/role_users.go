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

func NewRoleUsers() *RoleUsers {
	return &RoleUsers{}
}

type RoleUsers struct {
}

func (s RoleUsers) GetList(c *ginx.Context) {
	d := conn.NewRoleUsersData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("RoleUsers GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s RoleUsers) Add(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleUsers Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleUsers{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleUsers Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("RoleUsers Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RoleUsers) Edit(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleUsers Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleUsers{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleUsers Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("RoleUsers Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RoleUsers) Delete(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleUsers Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleUsers{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RoleUsers Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("RoleUsers Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
