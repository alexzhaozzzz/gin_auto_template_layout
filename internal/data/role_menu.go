package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleMenuData() *RoleMenuData {
	return &RoleMenuData{}
}

type RoleMenuData struct {
}

func (s RoleMenuData) List() ([]*po.SysRoleMenu, error) {
	list := make([]*po.SysRoleMenu, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleMenuData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleMenuData) InfoByName(name string) (*po.SysRoleMenu, error) {
	info := &po.SysRoleMenu{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("RoleMenuData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleMenuData) Info(id int64) (*po.SysRoleMenu, error) {
	info := &po.SysRoleMenu{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleMenuData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleMenuData) Edit(eUser *po.SysRoleMenu) error {
	err := GetMerchantDB().Model(&po.SysRoleMenu{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleMenuData) Delete(eUser *po.SysRoleMenu) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
