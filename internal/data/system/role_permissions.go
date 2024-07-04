package system

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewRolePermissionsData() *RolePermissionsData {
	return &RolePermissionsData{}
}

type RolePermissionsData struct {
}

func (s RolePermissionsData) ListByRoleId(roleId int64) ([]*po.SysRolePermissions, error) {
	list := make([]*po.SysRolePermissions, 0)
	err := conn.GetMerchantDB().Where("role_id = ?", roleId).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RolePermissionsData) Info(id int64) (*po.SysRolePermissions, error) {
	info := &po.SysRolePermissions{}
	err := conn.GetMerchantDB().First(info, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RolePermissionsData) Add(rp *po.SysRolePermissions) error {
	rp.Id = 0
	err := conn.GetMerchantDB().Model(&po.SysRolePermissions{}).Create(rp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) Edit(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Model(&po.SysRolePermissions{}).Where("id = ?", rp.Id).Updates(rp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) DeleteByRoleId(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Where("role_id = ?", rp.RoleId).Delete(rp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RolePermissionsData) DeleteByRoleIdAndPermissionId(rp *po.SysRolePermissions) error {
	err := conn.GetMerchantDB().Where("role_id = ? AND permission_id = ?", rp.RoleId, rp.PermissionId).Delete(rp).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RolePermissionsData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
