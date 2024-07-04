package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
)

func NewBetStatisticsData() *BetStatisticsData {
	return &BetStatisticsData{}
}

type BetStatisticsData struct {
}

func (s BetStatisticsData) ListByPage(cIds []int32, page *dto.UserBetReq) ([]*po.BetStatistics, error) {
	list := make([]*po.BetStatistics, 0)
	err := conn.GetLogDB().Where("channel_id IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("BetStatisticsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s BetStatisticsData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetLogDB().Model(&po.BetStatistics{}).Where("channel_id IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("BetStatisticsData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s BetStatisticsData) InfoById(id uint64) (*po.BetStatistics, error) {
	info := &po.BetStatistics{}
	err := conn.GetLogDB().First(info, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("BetStatisticsData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
