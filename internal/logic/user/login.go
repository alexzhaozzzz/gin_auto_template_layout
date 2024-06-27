package user

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/user"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewLogin() *Login {
	return &Login{}
}

type Login struct {
}

func (s Login) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Login GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.Login{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Login GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Login GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	//if req.Day == "" {
	//	now := time.Now()
	//	req.Day = fmt.Sprintf("%d%d%d", now.YearDay(), now.Month(), now.Day())
	//}

	d := user.NewLoginData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Login GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("Login GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": list, "page": pageResp}
	c.RenderSuccess(resp)
	return
}
