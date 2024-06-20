package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleUsersData() *RoleUsersData {
	return &RoleUsersData{}
}

type RoleUsersData struct {
}

func (s RoleUsersData) List() ([]*po.SysRoleUsers, error) {
	list := make([]*po.SysRoleUsers, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleUsersData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleUsersData) InfoByName(name string) (*po.SysRoleUsers, error) {
	info := &po.SysRoleUsers{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("RoleUsersData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleUsersData) Info(id int64) (*po.SysRoleUsers, error) {
	info := &po.SysRoleUsers{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleUsersData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleUsersData) Edit(eUser *po.SysRoleUsers) error {
	err := GetMerchantDB().Model(&po.SysRoleUsers{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleUsersData) Delete(eUser *po.SysRoleUsers) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
