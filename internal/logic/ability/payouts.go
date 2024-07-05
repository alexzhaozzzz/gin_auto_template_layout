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

func NewPayouts() *Payouts {
	return &Payouts{}
}

type Payouts struct {
}

// GetListByPage ...
// @Summary 提现日志
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Param order_id query string false "订单id"
// @Param state query int false "订单状态: 0=审核中, 1=审核成功,2=提现成功,-1=审核失败,-2=提现失败,-3=提现提交失败"
// @Param order_start_time query int false "订单开始时间（时间戳）"
// @Param order_end_time query int false "订单结束时间（时间戳）"
// @Param loan_start_time query int false "放款开始时间（时间戳）"
// @Param loan_end_time query int false "放款结束时间（时间戳）"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.WithdrawRecord,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/payouts [get]
func (s Payouts) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.PayoutsReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ListCount Err: %s", err.Error())
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

// GetAuditListByPage ...
// @Summary 提现审核
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Param order_id query string false "订单id"
// @Param amount_start query int false "金额开始值"
// @Param amount_end query int false "金额结束值"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.PayoutsAuditResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/payoutsaudit [get]
func (s Payouts) GetAuditListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.PayoutsAuditReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByStatePage(cIds, 0, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	respList := make([]dto.PayoutsAuditResp, 0)
	if len(list) > 0 {

		dpp := ability.NewPlayerProfileData()

		for _, v := range list {
			profileInfo, err := dpp.InfoById(v.Playerid)
			if err != nil {
				continue
			}

			respList = append(respList, dto.PayoutsAuditResp{
				Id:            v.Id,
				Orderid:       v.Orderid,
				Playerid:      v.Playerid,
				Amount:        v.Amount,
				State:         v.State,
				Des:           v.Des,
				Createtime:    v.Createtime.Format(time.DateTime),
				Paytime:       v.Paytime.Format(time.DateTime),
				Checktime:     v.Checktime.Format(time.DateTime),
				Coin:          v.Coin,
				Channel:       v.Channel,
				ExOrderid:     v.ExOrderid,
				Bankinfo:      v.Bankinfo,
				ExResp:        v.ExResp,
				BackCoin:      v.BackCoin,
				Operator:      v.Operator,
				Fee:           v.Fee,
				From:          v.From,
				Partnerid:     v.Partnerid,
				Platform:      v.Platform,
				After:         v.After,
				Kind:          v.Kind,
				ShowErr:       v.ShowErr,
				CardId:        v.CardId,
				PayCreatetime: v.PayCreatetime,
				Retry:         v.Retry,
				RetryState:    v.RetryState,
				QueryState:    v.QueryState,
				QueryResp:     v.QueryResp,
				LastOrderId:   v.LastOrderId,
				Red:           v.Red,
				WdFlag:        v.WdFlag,
				UserFlag:      v.UserFlag,
				UserKind:      profileInfo.NKind,
			})
		}
	}

	total, err := d.ListStateCount(cIds, 0, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetListByPage ListCount Err: %s", err.Error())
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

// GetUserBaseInfo ...
// @Summary 用户基本信息
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Success 200 {object} ginx.Result{data=ginx.InfoResponses{info=dto.UserBaseResp}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/userbaseinfo [get]
func (s Payouts) GetUserBaseInfo(c *ginx.Context) {
	req := &dto.UserBaseReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetUserFoundationInfo ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dp := ability.NewPlayerData()
	dpInfo, err := dp.InfoById(req.PlayerId)
	if err != nil {
		logrus.Errorf("PayoutsReq dp.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpd := ability.NewPlayerDeviceData()
	dpdInfo, err := dpd.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpd.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpp := ability.NewPlayerProfileData()
	dppInfo, err := dpp.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpp.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dpb := ability.NewPlayerBankData()
	dpbInfo, err := dpb.InfoById(uint64(req.PlayerId))
	if err != nil {
		logrus.Errorf("PayoutsReq dpb.InfoById Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	respInfo := dto.UserBaseResp{
		PlayerId:   req.PlayerId,
		Tag:        dppInfo.NTag,
		UType:      dppInfo.NType,
		NickName:   dpInfo.Nickname,
		Card:       dpbInfo.NBankcardid,
		CardName:   dpbInfo.NName,
		Phone:      dpbInfo.NPhone,
		DeviceCode: dpdInfo.NDevCode,
	}

	resp := ginx.InfoResponses{
		Info: respInfo,
	}
	c.RenderSuccess(resp)
	return
}

// GetUserRechargeListByPage ...
// @Summary 用户充值提现
// @Tags Ability
// @Produce  json
// @Param player_id query int false "玩家id"
// @Param order_start_time query int false "订单开始时间（时间戳）"
// @Param order_end_time query int false "订单结束时间（时间戳）"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.WithdrawRecord,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/userrechargelog [get]
func (s Payouts) GetUserRechargeListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("PayoutsReq GetUserRechargeListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.UserRechargeListReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("PayoutsReq GetUserRechargeListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("PayoutsReq GetUserRechargeListByPage GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewPayoutsData()
	list, err := d.ListByPlayerPage(cIds, req.PlayerId, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetUserRechargeListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListPlayerCount(cIds, req.PlayerId, req)
	if err != nil {
		logrus.Errorf("PayoutsReq GetUserRechargeListByPage ListCount Err: %s", err.Error())
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
