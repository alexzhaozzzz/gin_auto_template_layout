package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewFirstPayData() *FirstPayData {
	return &FirstPayData{}
}

type FirstPayData struct {
}

//func (s FirstPayData) ListByPage(cIds []int32, page *dto.LoginLog) ([]*po.PayFirstRecord, error) {
//	list := make([]*po.PayFirstRecord, 0)
//	err := conn.GetPlayerDB().Where("n_channel IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		logrus.Errorf("FirstPayData List Err: %s", err.Error())
//		return nil, err
//	}
//	return list, nil
//}
//
//func (s FirstPayData) ListCount(cIds []int32) (int64, error) {
//	num := int64(0)
//	err := conn.GetPlayerDB().Model(&po.PayFirstRecord{}).Where("n_channel IN ?", cIds).Count(&num).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		logrus.Errorf("FirstPayData List Err: %s", err.Error())
//		return num, err
//	}
//	return num, nil
//}

func (s FirstPayData) ListByUids(ids []int64) ([]po.PayFirstRecord, error) {
	list := make([]po.PayFirstRecord, 0)
	err := conn.GetPlayerDB().Find(&list, "id IN ?", ids).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("MerchantData Info Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s FirstPayData) Info(uid int64) (*po.PayFirstRecord, error) {
	info := &po.PayFirstRecord{}
	err := conn.GetPlayerDB().First(info, "n_uid = ?", uid).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("MerchantData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
