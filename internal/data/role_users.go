package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleUsersData() *RoleUsersData {
	return &RoleUsersData{}
}

type RoleUsersData struct {
}

func (s RoleUsersData) List() ([]*po.SysRoleUsers, error) {
	list := make([]*po.SysRoleUsers, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleUsersData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleUsersData) Info(id int64) (*po.SysRoleUsers, error) {
	info := &po.SysRoleUsers{}
	err := conn.GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleUsersData) Add(ru *po.SysRoleUsers) error {
	err := conn.GetMerchantDB().Model(&po.SysRoleUsers{}).Where("id = ?", &ru.Id).Create(&ru).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleUsersData) Edit(ru *po.SysRoleUsers) error {
	err := conn.GetMerchantDB().Model(&po.User{}).Where("id = ?", &ru.Id).Updates(ru).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleUsersData) Delete(ru *po.SysRoleUsers) error {
	err := conn.GetMerchantDB().Where("id = ?", &ru.Id).Delete(&ru).Error
	if err != nil {
		logrus.Errorf("RoleUsersData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
