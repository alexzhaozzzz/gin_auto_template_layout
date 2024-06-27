package user

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewGoldChangesData() *GoldChangesData {
	return &GoldChangesData{}
}

type GoldChangesData struct {
}

func (s GoldChangesData) GetTableName(day string) string {
	cr := po.CurrencyRecord{}
	table := fmt.Sprintf("%s_%s", cr.TableName(), day)
	return table
}

func (s GoldChangesData) ListByPage(day string, cIds []int32, page *dto.CurrencyRecord) ([]*po.CurrencyRecord, error) {
	list := make([]*po.CurrencyRecord, 0)
	err := conn.GetLogDB().Table(s.GetTableName(day)).Where("channel IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("GoldChangesData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s GoldChangesData) ListCount(day string, cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetLogDB().Table(s.GetTableName(day)).Model(&po.CurrencyRecord{}).Where("channel IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("GoldChangesData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
