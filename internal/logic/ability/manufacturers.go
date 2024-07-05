package ability

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
)

func NewManufacturers() *Manufacturers {
	return &Manufacturers{}
}

type Manufacturers struct {
}

// GetListByPage ...
// @Summary 厂商统计
// @Tags Statistics
// @Produce  json
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.ManufacturersResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /statistics/manufacturers [get]
func (s Manufacturers) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Manufacturers GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.StatisticalLogPlatformReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Manufacturers GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Manufacturers GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewManufacturersData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Manufacturers GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]dto.ManufacturersResp, 0)
	if len(list) > 0 {
		for _, v := range list {

			var rtp float64
			if v.Reward > 0 && v.Bet > 0 {
				rtp = float64(v.Reward / v.Bet)
			}

			respList = append(respList, dto.ManufacturersResp{
				Win:       v.Win,
				Id:        v.Id,
				ChannelId: v.ChannelId,
				Reward:    v.Reward,
				Date:      v.Date,
				Bet:       v.Bet,
				Chip:      v.Chip,
				JoinNum:   v.JoinNum,
				Platform:  v.Platform,
				Rtp:       rtp,
			})
		}
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("Manufacturers GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := ginx.ListResponses{
		List: respList,
		Page: pageResp,
	}
	c.RenderSuccess(resp)
	return
}
