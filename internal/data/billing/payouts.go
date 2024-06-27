package billing

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewPayoutsData() *PayoutsData {
	return &PayoutsData{}
}

type PayoutsData struct {
}

func (s PayoutsData) ListByPage(cIds []int32, page *dto.Payouts) ([]*po.WithdrawRecord, error) {
	list := make([]*po.WithdrawRecord, 0)
	err := conn.GetPlayerDB().Where("n_channel IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PayoutsData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.WithdrawRecord{}).Where("n_channel IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PayoutsData) ListByStatePage(cIds []int32, state int32, page *dto.Payouts) ([]*po.WithdrawRecord, error) {
	list := make([]*po.WithdrawRecord, 0)
	err := conn.GetPlayerDB().Where("n_channel IN ? AND n_state = ? ", cIds, state).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PayoutsData) ListStateCount(cIds []int32, state int32) (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.WithdrawRecord{}).Where("n_channel IN ? AND n_state = ? ", cIds, state).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
