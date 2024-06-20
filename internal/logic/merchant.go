package logic

import (
	"github.com/sirupsen/logrus"
	podata "gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewMerchant() *Merchant {
	return &Merchant{}
}

type Merchant struct {
}

func (s Merchant) GetList(c *ginx.Context) {
	d := podata.NewMerchantData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Merchant GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	data := map[string]interface{}{"list": list}
	c.RenderSuccess(data)
	return
}
