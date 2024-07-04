package system

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/util"
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

	d := system.NewUserData()
	info, err := d.InfoByName(req.Username)
	if err != nil {
		logrus.Errorf("User Login get Info Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if info.Password != util.MD5(req.Password) {
		logrus.Errorf("User Login Password Err")
		c.Render(statusx.StatusPasswordError, nil)
		return
	}

	dru := system.NewRoleUsersData()
	ruInfo, err := dru.InfoByUserId(info.Id)
	if err != nil {
		logrus.Errorf("User InfoByUserId Err: %s", err.Error())
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
	d := system.NewUserData()
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

func (s User) GetListByPage(c *ginx.Context) {
	req := &dto.Pagination{}
	if err := c.ShouldBindQuery(req); err != nil {
		logrus.Errorf("User GetListByPage ShouldBindQuery Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	d := system.NewUserData()
	list, err := d.ListByPage(req)
	if err != nil {
		logrus.Errorf("User GetListByPage Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	total, err := d.ListCount()
	if err != nil {
		logrus.Errorf("User GetListByPage ListCount Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	req.TotalNum = int(total)

	resp := map[string]interface{}{"list": list, "page": req}
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

	d := system.NewUserData()
	info, err := d.Info(int64(idInt))
	if err != nil {
		logrus.Errorf("User GetInfo Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	dru := system.NewRoleUsersData()
	roleInfo, err := dru.InfoByUid(int64(idInt))
	if err != nil {
		logrus.Errorf("User InfoByUid Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	resp := map[string]interface{}{"info": info, "role": roleInfo}
	c.RenderSuccess(resp)
	return
}

func (s User) Add(c *ginx.Context) {
	req := &dto.UserChangeReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Add ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	if req.User.Username == "" || req.User.Password == "" {
		logrus.Errorf("User Add Param Err")
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.User{
		Username:   req.User.Username,
		Password:   util.MD5(req.User.Password),
		MerchantId: req.User.MerchantId,
		NickName:   req.User.NickName,
		Phone:      req.User.Phone,
		Email:      req.User.Email,
		Avatar:     req.User.Avatar,
	}
	d := system.NewUserData()
	uid, err := d.Add(&dReq)
	if err != nil {
		logrus.Errorf("User Add Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}
	if uid <= 0 {
		logrus.Errorf("User Add Db Uid Err")
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.RoleChange(req.RoleUser.RoleId, uid)
	if err != nil {
		logrus.Errorf("User Add RoleChange Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s User) Edit(c *ginx.Context) {
	req := &dto.UserChangeReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Errorf("User Edit ShouldBindJSON Err: %s", err.Error())
		c.Render(statusx.StatusInvalidRequest, nil)
		return
	}

	dReq := po.User{
		Id:         req.User.Id,
		Username:   req.User.Username,
		Password:   util.MD5(req.User.Password),
		MerchantId: req.User.MerchantId,
		NickName:   req.User.NickName,
		Phone:      req.User.Phone,
		Email:      req.User.Email,
		Avatar:     req.User.Avatar,
	}
	d := system.NewUserData()
	err := d.Edit(&dReq)
	if err != nil {
		logrus.Errorf("User Edit Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.RoleChange(req.RoleUser.RoleId, req.User.Id)
	if err != nil {
		logrus.Errorf("User Edit RoleChange Err: %s", err.Error())
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

	dReq := po.User{
		Id: req.Id,
	}
	d := system.NewUserData()
	err := d.Delete(&dReq)
	if err != nil {
		logrus.Errorf("User Delete Db Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	err = s.RoleChange(0, req.Id)
	if err != nil {
		logrus.Errorf("User Edit RoleChange Err: %s", err.Error())
		c.Render(statusx.StatusInternalServerError, nil)
		return
	}

	c.RenderSuccess(nil)
	return
}

func (s User) RoleChange(roleId int64, userId int64) error {
	d := system.NewRoleUsersData()
	delReq := po.SysRoleUsers{
		UserId: userId,
	}
	err := d.DeleteByUserId(&delReq)
	if err != nil {
		logrus.Errorf("RoleChange DeleteByRoleId Err: %s", err.Error())
		return errors.New("RoleChange DeleteByRoleId Err")
	}

	if roleId > 0 {
		addReq := po.SysRoleUsers{
			RoleId: roleId,
			UserId: userId,
		}
		err = d.Add(&addReq)
		if err != nil {
			logrus.Errorf("RoleChange Add To Db Err: %s", err.Error())
			return errors.New("RoleChange Add To Db Err")
		}
	}

	return nil
}
