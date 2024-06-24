package logic

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewMenu() *Menu {
	return &Menu{}
}

type Menu struct {
}

func (s Menu) GetList(c *ginx.Context) {
	d := data.NewMenuData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("Menu GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	reqList := make([]*dto.SysMenu, 0)
	if len(list) > 0 {
		for _, v := range list {
			pIds := make([]int64, 0)
			if v.PermissionsIds != "" {
				_ = json.Unmarshal([]byte(v.PermissionsIds), &pIds)
			}

			reqList = append(reqList, &dto.SysMenu{
				Id:             v.Id,
				Title:          v.Title,
				Icon:           v.Icon,
				Sort:           v.Sort,
				Show:           v.Show,
				ParentId:       v.ParentId,
				Uri:            v.Uri,
				PermissionsIds: pIds,
			})
		}
	}

	resp := map[string]interface{}{"list": reqList}
	c.RenderSuccess(resp)
	return
}

func (s Menu) GetInfo(c *ginx.Context) {
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

	d := data.NewMenuData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("Menu GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	pIds := make([]int64, 0)
	if info.PermissionsIds != "" {
		_ = json.Unmarshal([]byte(info.PermissionsIds), &pIds)
	}
	reqInfo := dto.SysMenu{
		Id:             info.Id,
		Title:          info.Title,
		Icon:           info.Icon,
		Sort:           info.Sort,
		Show:           info.Show,
		ParentId:       info.ParentId,
		Uri:            info.Uri,
		PermissionsIds: pIds,
	}

	resp := map[string]interface{}{"info": reqInfo}
	c.RenderSuccess(resp)
	return
}

func (s Menu) Add(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	permStr, err := json.Marshal(req.PermissionsIds)
	if err != nil {
		logrus.Errorf("Menu Add Json Marshal Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	dReq := po.SysMenu{
		Id:             req.Id,
		Title:          req.Title,
		Icon:           req.Icon,
		Sort:           req.Sort,
		Show:           req.Show,
		ParentId:       req.ParentId,
		Uri:            req.Uri,
		PermissionsIds: string(permStr),
	}

	d := data.NewMenuData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("Menu Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Menu) Edit(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	permStr, err := json.Marshal(req.PermissionsIds)
	if err != nil {
		logrus.Errorf("Menu Add Json Marshal Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}
	dReq := po.SysMenu{
		Id:             req.Id,
		Title:          req.Title,
		Icon:           req.Icon,
		Sort:           req.Sort,
		Show:           req.Show,
		ParentId:       req.ParentId,
		Uri:            req.Uri,
		PermissionsIds: string(permStr),
	}

	d := data.NewMenuData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("Menu Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s Menu) Delete(c *ginx.Context) {
	req := &dto.SysMenu{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("Menu Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.SysMenu{
		Id: req.Id,
	}
	d := data.NewMenuData()
	err := d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("Menu Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
