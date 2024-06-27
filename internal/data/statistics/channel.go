package statistics

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewChannelData() *ChannelData {
	return &ChannelData{}
}

type ChannelData struct {
}

func (s ChannelData) ListByPage(cIds []int32, page *dto.StatisticalDaily) ([]*po.StatisticalDaily, error) {
	list := make([]*po.StatisticalDaily, 0)
	err := conn.GetAdminDB().Where("channel_id IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s ChannelData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetAdminDB().Model(&po.StatisticalDaily{}).Where("channel_id IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
