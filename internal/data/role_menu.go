package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleMenuData() *RoleMenuData {
	return &RoleMenuData{}
}

type RoleMenuData struct {
}

func (s RoleMenuData) ListByRoleId(roleId int64) ([]*po.SysRoleMenu, error) {
	list := make([]*po.SysRoleMenu, 0)
	err := conn.GetMerchantDB().Where("role_id = ?", roleId).Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleMenuData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleMenuData) Info(id int64) (*po.SysRoleMenu, error) {
	info := &po.SysRoleMenu{}
	err := conn.GetMerchantDB().First(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleMenuData) Add(rm *po.SysRoleMenu) error {
	rm.Id = 0
	err := conn.GetMerchantDB().Model(&po.SysRoleMenu{}).Create(rm).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleMenuData) Edit(rm *po.SysRoleMenu) error {
	err := conn.GetMerchantDB().Model(&po.SysRoleMenu{}).Where("id = ?", rm.Id).Updates(rm).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleMenuData) DeleteByRoleId(rm *po.SysRoleMenu) error {
	err := conn.GetMerchantDB().Where("role_id = ?", rm.RoleId).Delete(rm).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Delete Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleMenuData) DeleteByRoleAndMenuId(rm *po.SysRoleMenu) error {
	err := conn.GetMerchantDB().Where("role_id = ? AND MenuId = ?", rm.RoleId, rm.MenuId).Delete(rm).Error
	if err != nil {
		logrus.Errorf("RoleMenuData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
