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

func NewRolePermissions() *RolePermissions {
	return &RolePermissions{}
}

type RolePermissions struct {
}

func (s RolePermissions) GetList(c *ginx.Context) {
	d := conn.NewRolePermissionsData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("RolePermissions GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s RolePermissions) Add(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RolePermissions Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRolePermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RolePermissions Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("RolePermissions Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RolePermissions) Edit(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RolePermissions Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRolePermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RolePermissions Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("RolePermissions Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s RolePermissions) Delete(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RolePermissions Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRolePermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("RolePermissions Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("RolePermissions Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
