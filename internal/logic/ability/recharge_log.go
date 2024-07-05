package ability

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ability"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/common"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"time"
)

func NewRechargeLog() *RechargeLog {
	return &RechargeLog{}
}

type RechargeLog struct {
}

// GetListByPage ...
// @Summary 充值日志
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Param order_id query string false "订单id"
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param finish_start_time query int false "充值结束开始时间（时间戳）"
// @Param finish_end_time query int false "充值结束结束时间（时间戳）"
// @Param amount query string false "金额"
// @Param state query int false "订单状态: 0=未付款, 1=已付款,2=已发货，3=补发成功， -1=充值失败，-2=发货失败"
// @Param platform query string false "充值平台"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.RechargeLogResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/rechargelog [get]
func (s RechargeLog) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("RechargeLog GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.RechargeLogReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("RechargeLog GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("RechargeLog GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewRechargeLogData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("RechargeLog GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]*dto.RechargeLogResp, 0)
	if len(list) > 0 {
		for _, v := range list {
			respList = append(respList, &dto.RechargeLogResp{
				Id:             v.Id,
				Orderid:        v.Orderid,
				Playerid:       v.Playerid,
				Amount:         v.Amount,
				State:          v.State,
				Des:            v.Des,
				Createtime:     v.Createtime.Format(time.DateTime),
				Paytime:        v.Paytime.Format(time.DateTime),
				Finishtime:     v.Finishtime.Format(time.DateTime),
				Coin:           v.Coin,
				Channel:        v.Channel,
				ExOrderid:      v.ExOrderid,
				Bankinfo:       v.Bankinfo,
				ExResp:         v.ExResp,
				From:           v.From,
				Phone:          v.Phone,
				Name:           v.Name,
				Partnerid:      v.Partnerid,
				Platform:       v.Platform,
				GoodsId:        v.GoodsId,
				Pic:            v.Pic,
				Inviteid:       v.Inviteid,
				Topinvite:      v.Topinvite,
				Src:            v.Src,
				InputType:      v.InputType,
				ClientType:     v.ClientType,
				ClientVer:      v.ClientVer,
				Entertime:      v.Entertime.Format(time.DateTime),
				BeforeCoin:     v.BeforeCoin,
				OpenSrc:        v.OpenSrc,
				VpayCreatetime: v.VpayCreatetime,
				Red:            v.Red,
			})
		}
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("RechargeLog GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	pageResp := &dto.Pagination{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		TotalNum:  int(total),
	}

	resp := ginx.ListResponses{
		List: list,
		Page: pageResp,
	}
	c.RenderSuccess(resp)
	return
}
