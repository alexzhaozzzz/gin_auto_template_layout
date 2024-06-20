package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewMenuData() *MenuData {
	return &MenuData{}
}

type MenuData struct {
}

func (s MenuData) List() ([]*po.SysMenu, error) {
	list := make([]*po.SysMenu, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("MenuData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s MenuData) InfoByName(name string) (*po.SysMenu, error) {
	info := &po.SysMenu{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("MenuData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MenuData) Info(id int64) (*po.SysMenu, error) {
	info := &po.SysMenu{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("MenuData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MenuData) Edit(eUser *po.SysMenu) error {
	err := GetMerchantDB().Model(&po.SysMenu{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("MenuData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MenuData) Delete(eUser *po.SysMenu) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("MenuData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
