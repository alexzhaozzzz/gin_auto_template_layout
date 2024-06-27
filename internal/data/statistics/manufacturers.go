package statistics

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewManufacturersData() *ManufacturersData {
	return &ManufacturersData{}
}

type ManufacturersData struct {
}

func (s ManufacturersData) ListByPage(cIds []int32, page *dto.StatisticalLogPlatform) ([]*po.StatisticalLogPlatform, error) {
	list := make([]*po.StatisticalLogPlatform, 0)
	err := conn.GetAdminDB().Where("channel_id IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ManufacturersData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s ManufacturersData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetAdminDB().Model(&po.StatisticalLogPlatform{}).Where("channel_id IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ManufacturersData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
