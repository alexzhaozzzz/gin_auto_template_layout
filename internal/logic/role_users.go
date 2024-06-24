package logic

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewRoleUsers() *RoleUsers {
	return &RoleUsers{}
}

type RoleUsers struct {
}

func (s RoleUsers) GetList(c *ginx.Context) {
	roleId := c.Param("role_id")
	if roleId == "" {
		logrus.Errorf("Menu GetInfo Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		logrus.Errorf("Menu GetInfo Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewRoleUsersData()
	list, err := d.ListByRoleId(int64(roleIdInt))
	if err != nil {
		logrus.Errorf("RoleUsers GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"list": list}
	c.RenderSuccess(resp)
	return
}

func (s RoleUsers) Add(c *ginx.Context) {
	req := &dto.SysRoleUsersIds{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleUsers Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if len(req.UserIds) <= 0 {
		logrus.Errorf("RoleUsers Add ShouldBindJSON Err: len(req.UserIds) <= 0")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewRoleUsersData()
	delReq := po.SysRoleUsers{
		RoleId: req.RoleId,
	}
	err := d.DeleteByRoleId(&delReq)
	if err != nil {
		logrus.Errorf("RoleUsers Add DeleteByRoleId Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	for _, v := range req.UserIds {
		addReq := po.SysRoleUsers{
			RoleId: req.RoleId,
			UserId: v,
		}
		err = d.Add(&addReq)
		if err != nil {
			logrus.Errorf("RoleUsers Add To Db Err: %s", err.Error())
			continue
		}
	}

	c.RenderSuccess(nil)
	return
}

func (s RoleUsers) Edit(c *ginx.Context) {
	req := &dto.SysRoleUsersIds{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("RoleUsers Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if len(req.UserIds) <= 0 {
		logrus.Errorf("RoleUsers Edit ShouldBindJSON Err: len(req.UserIds) <= 0")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewRoleUsersData()
	delReq := po.SysRoleUsers{
		RoleId: req.RoleId,
	}
	err := d.DeleteByRoleId(&delReq)
	if err != nil {
		logrus.Errorf("RoleUsers Edit DeleteByRoleId Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	for _, v := range req.UserIds {
		addReq := po.SysRoleUsers{
			RoleId: req.RoleId,
			UserId: v,
		}
		err = d.Add(&addReq)
		if err != nil {
			logrus.Errorf("RoleUsers Edit To Db Err: %s", err.Error())
			continue
		}
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

	dReq := po.SysRoleUsers{
		RoleId: req.RoleId,
		UserId: req.UserId,
	}

	d := data.NewRoleUsersData()
	err := d.DeleteByRoleIdAndUserId(&dReq)
	if err != nil {
		logrus.Errorf("RoleUsers Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	//d := data.NewRoleUsersData()
	//err = d.Delete(&dReq)
	//if err != nil {
	//	logrus.Errorf("RoleUsers Delete Db Err: %s", err.Error())
	//	c.Render(statusx.StatusInternalServerError, nil)
	//	return
	//}

	c.RenderSuccess(nil)
	return
}
