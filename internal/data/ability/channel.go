package ability

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

func (s ChannelData) ListByPage(cIds []int32, page *dto.StatisticalDailyReq) ([]*po.StatisticalDaily, error) {
	list := make([]*po.StatisticalDaily, 0)

	adminDb := conn.GetAdminDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("time_int BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.NewReg > 0 {
		adminDb = adminDb.Where("register = ?", page.NewReg)
	}
	if page.Recharge > 0 {
		adminDb = adminDb.Where("first_recharge_total = ?", page.Recharge)
	}

	err := adminDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s ChannelData) ListCount(cIds []int32, page *dto.StatisticalDailyReq) (int64, error) {
	num := int64(0)

	adminDb := conn.GetAdminDB().Model(&po.StatisticalDaily{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("time_int BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.NewReg > 0 {
		adminDb = adminDb.Where("register = ?", page.NewReg)
	}
	if page.Recharge > 0 {
		adminDb = adminDb.Where("first_recharge_total = ?", page.Recharge)
	}

	err := adminDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s ChannelData) SlaveInfoById(id int64) (*po.StatisticalDailySlave, error) {
	info := &po.StatisticalDailySlave{}
	err := conn.GetAdminDB().First(info, "s_d_id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData Info Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
