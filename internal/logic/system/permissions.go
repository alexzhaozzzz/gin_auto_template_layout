package system

import (
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func NewPermissions() *Permissions {
	return &Permissions{}
}

type Permissions struct {
}

func (s Permissions) GetList(c *ginx.Context) {
	resp := map[string]interface{}{"list": bootstrap.PermConfigList}
	c.RenderSuccess(resp)
	return
}
