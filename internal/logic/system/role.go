package system

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewRole() *Role {
	return &Role{}
}

type Role struct {
}

func (s Role) GetList(c *ginx.Context) {
	d := system.NewRoleData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Role GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"list": list}
	c.RenderSuccess(resp)
	return
}

func (s Role) GetListByPage(c *ginx.Context) {
	req := &dto.Pagination{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Role GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewRoleData()
	list, err := d.ListByPage(req)
	if err != nil {
		logrus.Errorf("Role GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount()
	if err != nil {
		logrus.Errorf("Role GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	req.TotalNum = int(total)

	resp := map[string]interface{}{"list": list, "page": req}
	c.RenderSuccess(resp)
	return
}

func (s Role) GetInfo(c *ginx.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Errorf("Role GetInfo Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("Role GetInfo Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewRoleData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("Role GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"info": info}
	c.RenderSuccess(resp)
	return
}

func (s Role) Add(c *ginx.Context) {
	req := &dto.SysRolesChange{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoles{
		Id:         0,
		Guid:       req.Role.Guid,
		MerchantId: req.Role.MerchantId,
		Name:       req.Role.Name,
		Code:       req.Role.Code,
	}
	d := system.NewRoleData()
	err, roleId := d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Role Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	if roleId <= 0 {
		logrus.Errorf("Role Add Db Err: roleId <= 0")
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.MenuChange(roleId, &req.RoleMenu)
	if err != nil {
		logrus.Errorf("Role Add MenuChange Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.PermissionsChange(roleId, &req.RolePermissions)
	if err != nil {
		logrus.Errorf("Role Add PermissionsChange Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) Edit(c *ginx.Context) {
	req := &dto.SysRolesChange{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Role Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysRoles{
		Id:         req.Role.Id,
		Guid:       req.Role.Guid,
		MerchantId: req.Role.MerchantId,
		Name:       req.Role.Name,
		Code:       req.Role.Code,
	}
	d := system.NewRoleData()
	err := d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Role Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.MenuChange(req.Role.Id, &req.RoleMenu)
	if err != nil {
		logrus.Errorf("Role Edit MenuChange Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.PermissionsChange(req.Role.Id, &req.RolePermissions)
	if err != nil {
		logrus.Errorf("Role Edit PermissionsChange Err: %s", err.Error())
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

	dReq := po.SysRoles{
		Id: req.Id,
	}
	d := system.NewRoleData()
	err := d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Role Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	dp := system.NewRolePermissionsData()
	dpReq := po.SysRolePermissions{
		RoleId: req.Id,
	}
	err = dp.DeleteByRoleId(&dpReq)
	if err != nil {
		logrus.Errorf("Role Delete Permissions DeleteByRoleId Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	dm := system.NewRoleMenuData()
	dmReq := po.SysRoleMenu{
		RoleId: req.Id,
	}
	err = dm.DeleteByRoleId(&dmReq)
	if err != nil {
		logrus.Errorf("Role Delete Menu DeleteByRoleId Err: %s", err.Error())
		return
	}

	dc := system.NewCasbinRuleData()
	delDcReq := po.CasbinRule{
		V0: strconv.FormatInt(req.Id, 10),
	}
	err = dc.DeleteByV0(&delDcReq)
	if err != nil {
		logrus.Errorf("Role Delete MenuChange DeleteByV0 Err: %s", err.Error())
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Role) PermissionsChange(roleId int64, req *dto.RolePermissions) error {
	d := system.NewRolePermissionsData()
	delReq := po.SysRolePermissions{
		RoleId: roleId,
	}
	err := d.DeleteByRoleId(&delReq)
	if err != nil {
		logrus.Errorf("Role PermissionsChange DeleteByRoleId Err: %s", err.Error())
		return err
	}

	dc := system.NewCasbinRuleData()
	delDcReq := po.CasbinRule{
		V0: strconv.FormatInt(roleId, 10),
	}
	err = dc.DeleteByV0(&delDcReq)
	if err != nil {
		logrus.Errorf("Role PermissionsChange DeleteByV0 Err: %s", err.Error())
		return err
	}

	if len(req.PermissionIds) > 0 {
		permList := bootstrap.PermConfigList

		for _, v := range req.PermissionIds {
			addReq := po.SysRolePermissions{
				RoleId:       roleId,
				PermissionId: v,
			}
			err = d.Add(&addReq)
			if err != nil {
				logrus.WithFields(logrus.Fields{"param": addReq}).Errorf(err.Error())
				continue
			}

			for _, v1 := range permList {
				if v1.Id == v {
					addCrReq := po.CasbinRule{
						Ptype: "p",
						V0:    strconv.FormatInt(roleId, 10),
						V1:    v1.Path,
						V2:    v1.Method,
					}
					err = dc.Add(&addCrReq)
					if err != nil {
						logrus.Errorf("Role PermissionsChange CasbinRule Add Db Err: %s", err.Error())
						continue
					}
				}
			}
		}
	}

	return nil
}

func (s Role) MenuChange(roleId int64, req *dto.RoleMenu) error {
	d := system.NewRoleMenuData()
	delReq := po.SysRoleMenu{
		RoleId: roleId,
	}
	err := d.DeleteByRoleId(&delReq)
	if err != nil {
		logrus.Errorf("Role MenuChange DeleteByRoleId Err: %s", err.Error())
		return err
	}

	if len(req.MenuIds) > 0 {
		for _, v := range req.MenuIds {
			addReq := po.SysRoleMenu{
				RoleId: roleId,
				MenuId: v,
			}
			err = d.Add(&addReq)
			if err != nil {
				logrus.Errorf("Role MenuChange SysRoleMenu Add Db Err: %s", err.Error())
				continue
			}
		}
	}

	return nil
}
