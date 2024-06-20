package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRolePermissionsData() *RolePermissionsData {
	return &RolePermissionsData{}
}

type RolePermissionsData struct {
}

func (s RolePermissionsData) List() ([]*po.SysRolePermissions, error) {
	list := make([]*po.SysRolePermissions, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RolePermissionsData) Info(id int64) (*po.SysRolePermissions, error) {
	info := &po.SysRolePermissions{}
	err := conn.GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RolePermissionsData) Add(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Model(&po.SysRolePermissions{}).Where("id = ?", &rp.Id).Create(&rp).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) Edit(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &rp.Id).Updates(rp).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) Delete(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Where("id = ?", &rp.Id).Delete(&rp).Error
	if err != nil {
		logrus.Errorf("RolePermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
