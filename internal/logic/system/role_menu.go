package system

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{}
}

type RoleMenu struct {
}

func (s RoleMenu) GetList(c *ginx.Context) {
	roleId := c.Param("role_id")
	if roleId == "" {
		logrus.Errorf("RoleMenu GetList Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		logrus.Errorf("RoleMenu GetList Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewRoleMenuData()
	list, err := d.ListByRoleId(int64(roleIdInt))
	if err != nil {
		logrus.Errorf("RoleMenu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"list": list}
	c.RenderSuccess(resp)
	return
}
