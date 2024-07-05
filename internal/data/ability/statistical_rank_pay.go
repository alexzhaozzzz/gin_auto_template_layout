package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
)

func NewStatisticalRankPayData() *StatisticalRankPayData {
	return &StatisticalRankPayData{}
}

type StatisticalRankPayData struct {
}

func (s StatisticalRankPayData) ListByPage(cIds []int32, page *dto.UserFightReq) ([]*po.StatisticalRankPay, error) {
	list := make([]*po.StatisticalRankPay, 0)

	adminDb := conn.GetAdminDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb = adminDb.Where("time_int BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.PlayerId > 0 {
		adminDb = adminDb.Where("player_id = ?", page.PlayerId)
	}
	err := adminDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("StatisticalRankPayData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s StatisticalRankPayData) ListCount(cIds []int32, page *dto.UserFightReq) (int64, error) {
	num := int64(0)

	adminDb := conn.GetAdminDB().Model(&po.StatisticalRankPay{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb = adminDb.Where("time_int BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.PlayerId > 0 {
		adminDb = adminDb.Where("player_id = ?", page.PlayerId)
	}
	err := adminDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("StatisticalRankPayData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s StatisticalRankPayData) InfoById(id uint64) (*po.StatisticalRankPay, error) {
	info := &po.StatisticalRankPay{}
	err := conn.GetAdminDB().First(info, "id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("StatisticalRankPayData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
