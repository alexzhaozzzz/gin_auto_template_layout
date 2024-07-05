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

// GetList ...
// @Summary 获取权限列表
// @Tags System
// @Produce  json
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]bootstrap.PermissionConfig}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /permission [get]
func (s Permissions) GetList(c *ginx.Context) {
	resp := ginx.ListResponses{
		List: bootstrap.PermConfigList,
	}
	c.RenderSuccess(resp)
	return
}
