package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewPermissionsData() *PermissionsData {
	return &PermissionsData{}
}

type PermissionsData struct {
}

func (s PermissionsData) List() ([]*po.SysPermissions, error) {
	list := make([]*po.SysPermissions, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("PermissionsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PermissionsData) InfoByName(name string) (*po.SysPermissions, error) {
	info := &po.SysPermissions{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("PermissionsData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s PermissionsData) Info(id int64) (*po.SysPermissions, error) {
	info := &po.SysPermissions{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("PermissionsData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s PermissionsData) Edit(eUser *po.SysPermissions) error {
	err := GetMerchantDB().Model(&po.SysPermissions{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("PermissionsData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s PermissionsData) Delete(eUser *po.SysPermissions) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("PermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
