package system

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
)

func NewUserData() *UserData {
	return &UserData{}
}

type UserData struct {
}

func (s UserData) List() ([]*po.User, error) {
	list := make([]*po.User, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("UserData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s UserData) ListByPage(page *dto.Pagination) ([]*po.User, error) {
	list := make([]*po.User, 0)
	err := conn.GetMerchantDB().Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil {
		logrus.Errorf("UserData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s UserData) ListCount() (int64, error) {
	num := int64(0)
	err := conn.GetMerchantDB().Model(&po.User{}).Count(&num).Error
	if err != nil {
		logrus.Errorf("UserData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s UserData) Info(id int64) (*po.User, error) {
	info := &po.User{}
	err := conn.GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("UserData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s UserData) InfoByName(name string) (*po.User, error) {
	info := &po.User{}
	err := conn.GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("UserData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s UserData) Add(user *po.User) (int64, error) {
	user.Id = 0
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &user.Id).Create(&user).Error
	if err != nil {
		logrus.Errorf("UserData Add Err: %s", err.Error())
		return 0, err
	}
	return user.Id, nil
}

func (s UserData) Edit(user *po.User) error {
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &user.Id).Updates(user).Error
	if err != nil {
		logrus.Errorf("UserData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s UserData) Delete(user *po.User) error {
	err := conn.GetMerchantDB().Where("id = ?", &user.Id).Delete(&user).Error
	if err != nil {
		logrus.Errorf("UserData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
