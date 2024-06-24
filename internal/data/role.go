package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewRoleData() *RoleData {
	return &RoleData{}
}

type RoleData struct {
}

func (s RoleData) List() ([]*po.SysRoles, error) {
	list := make([]*po.SysRoles, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("RoleData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RoleData) Info(id int64) (*po.SysRoles, error) {
	info := &po.SysRoles{}
	err := conn.GetMerchantDB().First(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("RoleData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s RoleData) Add(ro *po.SysRoles) (error, int64) {
	ro.Id = 0
	err := conn.GetMerchantDB().Model(&po.SysRoles{}).Create(ro).Error
	if err != nil {
		logrus.Errorf("RoleData Add Err: %s", err.Error())
		return err, 0
	}
	return nil, ro.Id
}

func (s RoleData) Edit(ro *po.SysRoles) error {
	err := conn.GetMerchantDB().Model(&po.SysRoles{}).Where("id = ?", ro.Id).Updates(ro).Error
	if err != nil {
		logrus.Errorf("RoleData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s RoleData) Delete(ro *po.SysRoles) error {
	err := conn.GetMerchantDB().Where("id = ?", ro.Id).Delete(ro).Error
	if err != nil {
		logrus.Errorf("RoleData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
