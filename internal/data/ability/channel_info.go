package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewChannelInfoData() *ChannelInfoData {
	return &ChannelInfoData{}
}

type ChannelInfoData struct {
}

func (s ChannelInfoData) ListByPage(cIds []int32, page *dto.StatisticalChannelInfoReq) ([]*po.StatisticalChannelInfo, error) {
	list := make([]*po.StatisticalChannelInfo, 0)

	adminDB := conn.GetAdminDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDB.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	err := adminDB.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelInfoData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s ChannelInfoData) ListCount(cIds []int32, page *dto.StatisticalChannelInfoReq) (int64, error) {
	num := int64(0)

	adminDB := conn.GetAdminDB().Model(&po.StatisticalChannelInfo{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDB.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	err := adminDB.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelInfoData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
