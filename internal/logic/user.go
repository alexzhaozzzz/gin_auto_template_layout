package logic

import (
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"strconv"
)

func NewUser() *User {
	return &User{}
}

type User struct {
}

func (s User) Login(c *ginx.Context) {
	req := &dto.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Login ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewUserData()
	info, err := d.InfoByName(req.Username)
	if err != nil {
		logrus.Errorf("User Login get Info Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dru := data.NewRoleUsersData()
	ruInfo, err := dru.InfoByUserId(info.Id)
	if err != nil {
		logrus.Errorf("RoleUsers InfoByUserId Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	se := viper.GetString("jwt.secret")
	t := viper.GetInt32("jwt.token_expire")
	i := viper.GetString("jwt.issuer")
	su := viper.GetString("jwt.sub")
	ext := &auth.JwtExt{
		UId:        uint64(info.Id),
		MerchantId: uint64(info.MerchantId),
		RoleId:     uint64(ruInfo.RoleId),
	}
	token := auth.NewJwt(se, su, i, t, ext)

	resp := map[string]interface{}{"user": info, "token": token}
	c.RenderSuccess(resp)
	return
}

func (s User) GetList(c *ginx.Context) {
	d := data.NewUserData()
	list, err := d.List()
	if err != nil {
		logrus.Errorf("User GetList Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"list": list}
	c.RenderSuccess(resp)
	return
}

func (s User) GetInfo(c *ginx.Context) {
	id := c.Param("id")
	if id == "" {
		logrus.Errorf("User GetInfo Param Err: Id Is Empty")
		c.Render(statusx.StatusInvalidRequest, nil)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("User GetInfo Param Atoi Err: %s", err)
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewUserData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("User GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"info": info}
	c.RenderSuccess(resp)
	return
}

func (s User) Add(c *ginx.Context) {
	req := &dto.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.User{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("User Add copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewUserData()
	err = d.Add(&dReq)
	if err != nil {
		logrus.Errorf("User Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s User) Edit(c *ginx.Context) {
	req := &dto.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.User{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("User Edit copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewUserData()
	err = d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("User Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s User) Delete(c *ginx.Context) {
	req := &dto.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Delete ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.User{}
	err := copier.Copy(dReq, &req)
	if err != nil {
		logrus.Errorf("User Delete copier Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := data.NewUserData()
	err = d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("User Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}
