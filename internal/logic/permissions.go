package logic

import (
	"github.com/sirupsen/logrus"
	podata "gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewPermissions() *Permissions {
	return &Permissions{}
}

type Permissions struct {
}

func (s Permissions) GetList(c *ginx.Context) {
	d := podata.NewPermissionsData()
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
