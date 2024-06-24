package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewCasbinRuleData() *CasbinRuleData {
	return &CasbinRuleData{}
}

type CasbinRuleData struct {
}

func (s CasbinRuleData) List() ([]*po.CasbinRule, error) {
	list := make([]*po.CasbinRule, 0)
	err := conn.GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("CasbinRuleData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s CasbinRuleData) InfoByV0(roleId string) (*po.CasbinRule, error) {
	info := &po.CasbinRule{}
	err := conn.GetMerchantDB().First(info, "v0 = ?", roleId).Error
	if err != nil {
		logrus.Errorf("CasbinRuleData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s CasbinRuleData) Add(ro *po.CasbinRule) error {
	ro.Id = 0
	err := conn.GetMerchantDB().Model(&po.CasbinRule{}).Create(ro).Error
	if err != nil {
		logrus.Errorf("CasbinRuleData Add Err: %s", err.Error())
		return err
	}
	return nil
}

func (s CasbinRuleData) Edit(ro *po.CasbinRule) error {
	err := conn.GetMerchantDB().Model(&po.CasbinRule{}).Where("id = ?", ro.Id).Updates(ro).Error
	if err != nil {
		logrus.Errorf("CasbinRuleData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s CasbinRuleData) DeleteByV0(ro *po.CasbinRule) error {
	err := conn.GetMerchantDB().Where("v0 = ?", ro.V0).Delete(ro).Error
	if err != nil {
		logrus.Errorf("CasbinRuleData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
