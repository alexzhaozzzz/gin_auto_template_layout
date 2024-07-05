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

func NewLoginLog() *LoginLog {
	return &LoginLog{}
}

type LoginLog struct {
}

// GetListByPage ...
// @Summary 用户登录日志
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.LoginLogResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/loginlog [get]
func (s LoginLog) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("LoginLog GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.LoginLogReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("LoginLog GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("LoginLog GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewLoginLogData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("LoginLog GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	dfp := ability.NewFirstPayData()

	respList := make([]dto.LoginLogResp, 0)
	if len(list) > 0 {
		for _, v := range list {
			var fpt int64
			var fpn string
			var fpot int64
			var fpon string

			firstInfo, err := dfp.Info(v.Playerid)
			if err == nil {
				fpt = firstInfo.NCreateTime.Unix()
				fpn = firstInfo.NAmount
			}

			respList = append(respList, dto.LoginLogResp{
				FirstPayTime:     fpt,
				FirstPayNum:      fpn,
				Endtime:          v.Endtime,
				Osinfomation:     v.Osinfomation,
				Appid:            v.Appid,
				FirstPayoutsTime: fpot,
				FirstPayoutsNum:  fpon,
				Logindevice:      v.Logindevice,
				Begintime:        v.Begintime,
				Devicecode:       v.Devicecode,
				Nickname:         v.Nickname,
				Duration:         v.Duration,
				Location:         v.Location,
				Clientversion:    v.Clientversion,
				Playerid:         v.Playerid,
				Channelid:        v.Channelid,
				Type:             v.Type,
				Ip:               v.Ip,
			})
		}
	}

	total, err := d.ListCount(cIds)
	if err != nil {
		logrus.Errorf("LoginLog GetListByPage ListCount Err: %s", err.Error())
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
