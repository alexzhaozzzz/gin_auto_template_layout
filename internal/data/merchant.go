package data

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewMerchantData() *MerchantData {
	return &MerchantData{}
}

type MerchantData struct {
}

func (s MerchantData) List() ([]*po.MerchantList, error) {
	list := make([]*po.MerchantList, 0)
	err := GetMerchantDB().Find(&list).Error
	if err != nil {
		logrus.Errorf("MerchantData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s MerchantData) InfoByName(name string) (*po.MerchantList, error) {
	info := &po.MerchantList{}
	err := GetMerchantDB().Find(info, "username = ?", name).Error
	if err != nil {
		logrus.Errorf("MerchantData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MerchantData) Info(id int64) (*po.MerchantList, error) {
	info := &po.MerchantList{}
	err := GetMerchantDB().Find(info, "id = ?", id).Error
	if err != nil {
		logrus.Errorf("MerchantData info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}

func (s MerchantData) Edit(eUser *po.MerchantList) error {
	err := GetMerchantDB().Model(&po.MerchantList{}).Where("id = ?", eUser.Id).Updates(eUser).Error
	if err != nil {
		logrus.Errorf("MerchantData Edit Err: %s", err.Error())
		return err
	}
	return nil
}

func (s MerchantData) Delete(eUser *po.MerchantList) error {
	err := GetMerchantDB().Where("id = ?", eUser.Id).Delete(&eUser).Error
	if err != nil {
		logrus.Errorf("MerchantData Delete Err: %s", err.Error())
		return err
	}
	return nil
}
