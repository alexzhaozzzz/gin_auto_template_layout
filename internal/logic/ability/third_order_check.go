package ability

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"time"
)

func NewThirdOrderCheck() *ThirdOrderCheck {
	return &ThirdOrderCheck{}
}

type ThirdOrderCheck struct{}

// CheckOrderStatus ...
// @Summary 获取订单第三方状态
// @Tags Ability
// @Produce  json
// @Param order_id query string false "订单id"
// @Success 200 {object} ginx.Result{data=ginx.InfoResponses{info=dto.ToThirdOrderCheckResp}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /ability/checkorderstatus [get]
func (s ThirdOrderCheck) CheckOrderStatus(c *ginx.Context) {
	req := &dto.ThirdOrderCheckReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	jwtInfo, ok := auth.GetJwtExt(c)
	if !ok {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus GetJwtExt Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	if jwtInfo.MerchantId <= 0 {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus MerchantId Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	url := viper.GetString("url.third_order_check")
	if url == "" {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus Post Url Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	cReqData := dto.ToThirdOrderCheckData{
		OrderNo: req.OrderId,
		Pass:    "asdfsafas@13131safdafda",
		Desc:    "查询支付订单",
		User:    "admin",
	}

	cReq := dto.ToThirdOrderCheckReq{
		Server: "shop",
		Cmd:    "query_pay_order",
		Data:   cReqData,
	}

	cResp := dto.ToThirdOrderCheckResp{}

	// Create a Resty Client
	client := resty.New()
	// 设置超时时间为 5 秒钟
	client.SetTimeout(5 * time.Second)
	CliResp, err := client.R().
		SetBody(cReq).
		SetResult(&cResp). // or SetResult(AuthSuccess{}).
		Post(url)
	if err != nil {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus Resp Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	if !CliResp.IsSuccess() {
		logrus.Errorf("ThirdOrderCheck CheckOrderStatus Resp Is Not Success")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	resp := ginx.InfoResponses{
		Info: cResp,
	}
	c.RenderSuccess(resp)
	return
}
