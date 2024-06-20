package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRolePermissionsData() *RolePermissionsData {
	return &RolePermissionsData{}
}

type RolePermissionsData struct {
}

func (s RolePermissionsData) List() ([]*po.SysRolePermissions, error) {
	list := make([]*po.SysRolePermissions, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RolePermissionsData) InfoByName(name string) (*po.SysRolePermissions, error) {
	info := &po.SysRolePermissions{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RolePermissionsData) Info(id int64) (*po.SysRolePermissions, error) {
	info := &po.SysRolePermissions{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RolePermissionsData) Edit(eUser *po.SysRolePermissions) error {
	err := GetMerchantDB().Model(&po.SysRolePermissions{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) Delete(eUser *po.SysRolePermissions) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
