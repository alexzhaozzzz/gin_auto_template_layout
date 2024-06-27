package system

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
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
		logrus.Errorf("RoleUsers GetList Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		logrus.Errorf("RoleUsers GetList Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewRoleUsersData()
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

//func (s RoleUsers) Delete(c *ginx.Context) {
//	req := &dto.SysRoleUsers{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		logrus.Errorf("RoleUsers Delete ShouldBindJSON Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	dReq := po.SysRoleUsers{
//		RoleId: req.RoleId,
//		UserId: req.UserId,
//	}
//
//	d := data.NewRoleUsersData()
//	err := d.DeleteByRoleIdAndUserId(&dReq)
//	if err != nil {
//		logrus.Errorf("RoleUsers Delete Db Err: %s", err.Error())
//		c.Render(statusx.StatusInternalServerError, nil)
//		return
//	}
//
//	c.RenderSuccess(nil)
//	return
//}
