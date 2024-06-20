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

func NewPermissions() *Permissions {
	return &Permissions{}
}

type Permissions struct {
}

func (s Permissions) GetList(c *ginx.Context) {
	d := conn.NewPermissionsData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Permissions GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Permissions) Add(c *ginx.Context) {
	req := &dto.SysPermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Permissions Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysPermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Permissions Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewPermissionsData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Permissions Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Permissions) Edit(c *ginx.Context) {
	req := &dto.SysPermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Permissions Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysPermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Permissions Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewPermissionsData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Permissions Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Permissions) Delete(c *ginx.Context) {
	req := &dto.SysPermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Permissions Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysPermissions{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Permissions Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewPermissionsData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Permissions Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
