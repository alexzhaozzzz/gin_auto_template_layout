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
	err := conn.GetAdminDB().Where("channel_id IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("StatisticalRankPayData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s StatisticalRankPayData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetAdminDB().Model(&po.StatisticalRankPay{}).Where("channel_id IN ?", cIds).Count(&num).Error
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
