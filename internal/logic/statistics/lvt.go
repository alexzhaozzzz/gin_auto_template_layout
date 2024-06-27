package statistics

import (
	"encoding/json"
	"github.com/sirupsen/logrus"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/statistics"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewLvt() *Lvt {
	return &Lvt{}
}

type Lvt struct {
}

func (s Lvt) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Lvt GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.StatisticalPlayerLtv{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Lvt GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Lvt GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := statistics.NewLvtData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Lvt GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]dto.StatisticalPlayerLtvResp, 0)
	if len(list) > 0 {
		for _, v := range list {
			var dayContent dto.ContentInfo
			var dayDrawContent dto.ContentInfo
			var registerContent dto.ContentInfo

			if v.DayContent != "" {
				// 使用json.Unmarshal进行解码
				err := json.Unmarshal([]byte(v.DayContent), &dayContent)
				if err != nil {
				}
			}

			if v.DayContentWithdraw != "" {
				// 使用json.Unmarshal进行解码
				err := json.Unmarshal([]byte(v.DayContentWithdraw), &dayDrawContent)
				if err != nil {
				}
			}

			if v.RegisterContent != "" {
				// 使用json.Unmarshal进行解码
				err := json.Unmarshal([]byte(v.RegisterContent), &registerContent)
				if err != nil {
				}
			}

			respList = append(respList, dto.StatisticalPlayerLtvResp{
				Date:               v.Date,
				Type:               v.Type,
				Recharge:           v.Recharge,
				Withdraw:           v.Withdraw,
				Register:           v.Register,
				DayContent:         dayContent,
				DayContentWithdraw: dayDrawContent,
				RegisterContent:    registerContent,
			})
		}
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("Lvt GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := map[string]interface{}{"list": respList, "page": pageResp}
	c.RenderSuccess(resp)
	return
}
