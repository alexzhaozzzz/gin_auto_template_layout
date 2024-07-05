package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewBetStatisticsData() *BetStatisticsData {
	return &BetStatisticsData{}
}

type BetStatisticsData struct {
}

func (s BetStatisticsData) ListByPage(cIds []int32, page *dto.UserBetReq) ([]*po.BetStatistics, error) {
	list := make([]*po.BetStatistics, 0)

	logDb := conn.GetLogDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		logDb = logDb.Where("day BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.PlayerId > 0 {
		logDb = logDb.Where("player_id = ?", page.PlayerId)
	}
	err := logDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("BetStatisticsData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s BetStatisticsData) ListCount(cIds []int32, page *dto.UserBetReq) (int64, error) {
	num := int64(0)

	logDb := conn.GetLogDB().Model(&po.BetStatistics{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		logDb = logDb.Where("day BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.PlayerId > 0 {
		logDb = logDb.Where("player_id = ?", page.PlayerId)
	}
	err := logDb.Count(&num).Error
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
