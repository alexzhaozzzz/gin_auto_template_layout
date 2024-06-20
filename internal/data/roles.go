package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleData() *RoleData {
	return &RoleData{}
}

type RoleData struct {
}

func (s RoleData) List() ([]*po.SysRoles, error) {
	list := make([]*po.SysRoles, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleData) InfoByName(name string) (*po.SysRoles, error) {
	info := &po.SysRoles{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("RoleData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleData) Info(id int64) (*po.SysRoles, error) {
	info := &po.SysRoles{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleData) Add(eUser *po.SysRoles) error {
	err := GetMerchantDB().Model(&po.SysRoles{}).Create(eUser).Error
	if err != nil {
		logrus.Errorf("RoleData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleData) Edit(eUser *po.SysRoles) error {
	err := GetMerchantDB().Model(&po.SysRoles{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("RoleData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleData) Delete(eUser *po.SysRoles) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("RoleData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
