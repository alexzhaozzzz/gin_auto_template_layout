package system

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewMerchant() *Merchant {
	return &Merchant{}
}

type Merchant struct {
}

// GetListByPage ...
// @Summary 获取商户列表
// @Tags System
// @Produce  json
// @Success 200 {object} ginx.Result{data=ginx.ListResponses{list=[]po.MerchantList,page=dto.Pagination}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /merchant [get]
func (s Merchant) GetListByPage(c *ginx.Context) {
	req := &dto.Pagination{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("Merchant GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	list, err := d.ListByPage(req)
	if err != nil {
		logrus.Errorf("Merchant GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount()
	if err != nil {
		logrus.Errorf("Merchant GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	req.TotalNum = int(total)

	resp := ginx.ListResponses{
		List: list,
		Page: req,
	}
	c.RenderSuccess(resp)
	return
}

// GetInfo ...
// @Summary 获取商户详情
// @Tags System
// @Produce  json
// @Param id path int true "商户id"
// @Success 200 {object} ginx.Result{data=ginx.InfoResponses{info=po.MerchantList}} "成功"
// @Failure 400 {string} string "bad request"
// @Router /merchant/info/:id [get]
func (s Merchant) GetInfo(c *ginx.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Errorf("Menu GetInfo Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("Menu GetInfo Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("Merchant GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := ginx.InfoResponses{
		Info: info,
	}
	c.RenderSuccess(resp)
	return
}

// Add ...
// @Summary 商户新增
// @Tags System
// @Accept  json
// @Produce  json
// @Param info body dto.MerchantListReq true "商户信息"
// @Success 200 {object} ginx.Result "成功"
// @Failure 400 {string} string "bad request"
// @Router /merchant/add [post]
func (s Merchant) Add(c *ginx.Context) {
	req := &dto.MerchantListReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Merchant Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

// Edit ...
// @Summary 商户编辑
// @Tags System
// @Accept  json
// @Produce  json
// @Param info body dto.MerchantListReq true "商户信息"
// @Success 200 {object} ginx.Result "成功"
// @Failure 400 {string} string "bad request"
// @Router /merchant/edit [post]
func (s Merchant) Edit(c *ginx.Context) {
	req := &dto.MerchantListReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{}
	err := copier.Copy(&dReq, &req)
	if err != nil {
		logrus.Errorf("Merchant Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewMerchantData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

// Delete ...
// @Summary 商户删除
// @Tags System
// @Accept  json
// @Produce  json
// @Param info body dto.MerchantListReq true "商户信息（只需要传ID值）"
// @Success 200 {object} ginx.Result "成功"
// @Failure 400 {string} string "bad request"
// @Router /merchant/delete [post]
func (s Merchant) Delete(c *ginx.Context) {
	req := &dto.MerchantListReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Merchant Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.MerchantList{
		Id: int32(req.Id),
	}
	d := system.NewMerchantData()
	err := d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Merchant Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
