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

func NewChannel() *Channel {
	return &Channel{}
}

type Channel struct {
}

// GetListByPage ...
// @Summary 渠道分析
// @Tags Statistics
// @Produce  json
// @Param start_time query int false "开始时间（时间戳）"
// @Param end_time query int false "结束时间（时间戳）"
// @Param new_reg query int false "新增注册人数"
// @Param recharge query int false "累计充值"
// @Param page_index query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]dto.StatisticalDailyResp,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /statistics/channel [get]
func (s Channel) GetListByPage(c *ginx.Context) {
	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("Channel GetListByPage GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	req := &dto.StatisticalDailyReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Channel GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cIds, err := common.GetChannelIdsByMerchantId(jwtInfo.MerchantId)
	if err != nil {
		logrus.Errorf("Channel GetChannelIdsByMerchantId Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := ability.NewChannelData()
	list, err := d.ListByPage(cIds, req)
	if err != nil {
		logrus.Errorf("Channel GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	da := ability.NewAnalyticsData()

	respList := make([]dto.StatisticalDailyResp, 0)
	if len(list) > 0 {
		for _, v := range list {

			slaveInfo, err := d.SlaveInfoById(int64(v.Id))
			if err != nil {
				logrus.Errorf("Channel SlaveInfoById Err: %s", err.Error())
				continue
			}

			pu, err := da.InfoByChannelId(int64(v.ChannelId))
			if err != nil {
				logrus.Errorf("Channel InfoByChannelId Err: %s", err.Error())
				continue
			}

			respList = append(respList, dto.StatisticalDailyResp{
				Register:             v.Register,
				Login:                v.Login,
				Active:               v.Active,
				Prepaid:              v.Prepaid,
				PrepaidCount:         v.PrepaidCount,
				PrepaidNumber:        v.PrepaidNumber,
				Withdraw:             v.Withdraw,
				WithdrawNew:          v.WithdrawNew,
				WithdrawNumber:       v.WithdrawNumber,
				NewPay:               v.NewPay,
				Win:                  v.Win,
				Bet:                  v.Bet,
				BetRealTime:          v.BetRealTime,
				WinRealTime:          v.WinRealTime,
				Coins:                v.Coins,
				Type:                 v.Type,
				TimeInt:              v.TimeInt,
				OldPay:               v.OldPay,
				ReOldPay:             v.ReOldPay,
				FirstRecharge:        v.FirstRecharge,
				FirstRechargeTotal:   v.FirstRechargeTotal,
				ChannelId:            v.ChannelId,
				RegBlackList:         v.RegBlackList,
				ActiveAvgSpin:        v.ActiveAvgSpin,
				ActiveAvgBet:         v.ActiveAvgBet,
				TimeTag:              v.TimeTag,
				Installs:             v.Installs,
				FbInstalls:           v.FbInstalls,
				InstallsUpdateStatus: v.InstallsUpdateStatus,
				//RegisterContent:            slaveInfo.RegisterContent,
				//LoginContent:               slaveInfo.LoginContent,
				//GoldContent:                slaveInfo.GoldContent,
				SpinActive:      slaveInfo.SpinActive,
				Arpu:            slaveInfo.Arpu,
				NewArpu:         slaveInfo.NewArpu,
				OldArpu:         slaveInfo.OldArpu,
				OldArpuRecharge: slaveInfo.OldArpuRecharge,
				ActiveWithdraw:  slaveInfo.ActiveWithdraw,
				NewWithdraw:     slaveInfo.NewWithdraw,
				OldWithdraw:     slaveInfo.OldWithdraw,
				NoBindCoins:     slaveInfo.Coins,
				Ingots:          slaveInfo.Ingots,
				IngotsPay:       slaveInfo.IngotsPay,
				NBonusSum:       slaveInfo.NBonusSum,
				NewPayLimit:     slaveInfo.NewPayLimit,
				//RechargeContent:            slaveInfo.RechargeContent,
				//WithdrawContent:            slaveInfo.WithdrawContent,
				//FirstRechargeContent:       slaveInfo.FirstRechargeContent,
				Bankruptcy: slaveInfo.Bankruptcy,
				//NewPayContent:              slaveInfo.NewPayContent,
				//ReNewPayContent:            slaveInfo.ReNewPayContent,
				NewDeviceLoginNum:          slaveInfo.NewDeviceLoginNum,
				OldDeviceLoginNum:          slaveInfo.OldDeviceLoginNum,
				RechargeSuccessRate:        slaveInfo.RechargeSuccessRate,
				NewRechargeSuccessRate:     slaveInfo.NewRechargeSuccessRate,
				OldRechargeSuccessRate:     slaveInfo.OldRechargeSuccessRate,
				NewRechargeSuccessRateCurr: slaveInfo.NewRechargeSuccessRateCurr,
				OldRechargeSuccessRateCurr: slaveInfo.OldRechargeSuccessRateCurr,
				//InstallSrcContent:          slaveInfo.InstallSrcContent,
				//NotSdkContent:              slaveInfo.NotSdkContent,
				//ReOldPayContent:            slaveInfo.ReOldPayContent,
				Pv: pu.Pv,
				Uv: pu.Uv,
			})
		}
	}

	total, err := d.ListCount(cIds, req)
	if err != nil {
		logrus.Errorf("Channel GetListByPage ListCount Err: %s", err.Error())
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
