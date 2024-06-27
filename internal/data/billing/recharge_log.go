package billing

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewRechargeLogData() *RechargeLogData {
	return &RechargeLogData{}
}

type RechargeLogData struct {
}

func (s RechargeLogData) ListByPage(cIds []int32, page *dto.RechargeLog) ([]*po.RechargeRecord, error) {
	list := make([]*po.RechargeRecord, 0)
	err := conn.GetPlayerDB().Where("n_channel IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RechargeLogData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RechargeLogData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.RechargeRecord{}).Where("n_channel IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RechargeLogData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
