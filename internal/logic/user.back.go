package logic

//func NewUser() *User {
//	return &User{}
//}
//
//type User struct {
//}
//
//func (s User) Login(c *ginx.Context) {
//	req := &dto.User{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		logrus.Errorf("User Login ShouldBindJSON Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	d := conn.NewUserData()
//	info, err := d.InfoByName(req.Username)
//	if err != nil {
//		logrus.Errorf("User Login get Info Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	se := viper.GetString("jwt.secret")
//	t := viper.GetInt32("jwt.tokenExpire")
//	i := viper.GetString("jwt.issuer")
//	su := viper.GetString("jwt.sub")
//	ext := &auth.JwtExt{
//		UId:        uint64(info.Id),
//		MerchantId: uint64(info.MerchantId),
//	}
//	token := auth.NewJwt(se, su, i, t, ext)
//
//	data := map[string]interface{}{"user": info, "token": token}
//	c.RenderSuccess(data)
//	return
//}
//
//func (s User) GetList(c *ginx.Context) {
//	d := conn.NewUserData()
//	list, err := d.List()
//	if err != nil {
//		logrus.Errorf("User GetList Err: %s", err.Error())
//		c.Render(statusx.StatusInternalServerError, nil)
//		return
//	}
//
//	data := map[string]interface{}{"list": list}
//	c.RenderSuccess(data)
//	return
//}
//
//func (s User) Edit(c *ginx.Context) {
//	req := &dto.User{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		logrus.Errorf("User Edit ShouldBindJSON Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	pu := po.User{}
//	err := copier.Copy(pu, &req)
//	if err != nil {
//		logrus.Errorf("User Edit copier Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	d := conn.NewUserData()
//	err = d.Edit(&pu)
//	if err != nil {
//		logrus.Errorf("User Edit Db Err: %s", err.Error())
//		c.Render(statusx.StatusInternalServerError, nil)
//		return
//	}
//
//	c.RenderSuccess(nil)
//	return
//}
//
//func (s User) Delete(c *ginx.Context) {
//	req := &dto.User{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		logrus.Errorf("User Delete ShouldBindJSON Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	pu := po.User{}
//	err := copier.Copy(pu, &req)
//	if err != nil {
//		logrus.Errorf("User Delete copier Err: %s", err.Error())
//		c.Render(statusx.StatusInvalidRequest, nil)
//		return
//	}
//
//	d := conn.NewUserData()
//	err = d.Delete(&pu)
//	if err != nil {
//		logrus.Errorf("User Delete Db Err: %s", err.Error())
//		c.Render(statusx.StatusInternalServerError, nil)
//		return
//	}
//
//	c.RenderSuccess(nil)
//	return
//}
