package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewPermissionsData() *PermissionsData {
	return &PermissionsData{}
}

type PermissionsData struct {
}

func (s PermissionsData) List() ([]*po.SysPermissions, error) {
	list := make([]*po.SysPermissions, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("PermissionsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PermissionsData) Info(id int64) (*po.SysPermissions, error) {
	info := &po.SysPermissions{}
	err := conn.GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("PermissionsData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s PermissionsData) Add(perm *po.SysPermissions) error {
	err := conn.GetMerchantDB().Model(&po.SysPermissions{}).Where("id = ?", &perm.Id).Create(&perm).Error
	if err != nil {
		logrus.Errorf("PermissionsData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s PermissionsData) Edit(perm *po.SysPermissions) error {
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &perm.Id).Updates(perm).Error
	if err != nil {
		logrus.Errorf("PermissionsData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s PermissionsData) Delete(perm *po.SysPermissions) error {
	err := conn.GetMerchantDB().Where("id = ?", &perm.Id).Delete(&perm).Error
	if err != nil {
		logrus.Errorf("PermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
