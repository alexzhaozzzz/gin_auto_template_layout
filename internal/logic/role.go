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

func NewRole() *Role {
	return &Role{}
}

type Role struct {
}

func (s Role) GetList(c *ginx.Context) {
	d := conn.NewRoleData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Role GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) Add(c *ginx.Context) {
	req := &dto.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoles{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Role Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) Edit(c *ginx.Context) {
	req := &dto.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoles{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Role Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) Delete(c *ginx.Context) {
	req := &dto.SysRoles{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoles{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Role Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) GetMenuList(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetMenuList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Role GetMenuList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) EditMenu(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleMenu{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Role Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) DeleteMenu(c *ginx.Context) {
	req := &dto.SysRoleMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleMenu{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleMenuData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Role Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) GetUserList(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetUserList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Role GetUserList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) EditUser(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleUsers{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Role Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) DeleteUser(c *ginx.Context) {
	req := &dto.SysRoleUsers{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoleUsers{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRoleUsersData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Role Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) GetPermissionsList(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role GetPermissionsList ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Role GetPermissionsList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}

func (s Role) EditPermissions(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRolePermissions{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Role Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) DeletePermissions(c *ginx.Context) {
	req := &dto.SysRolePermissions{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRolePermissions{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("Role Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := conn.NewRolePermissionsData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Role Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
