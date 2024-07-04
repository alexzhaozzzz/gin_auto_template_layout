package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewLvtData() *LvtData {
	return &LvtData{}
}

type LvtData struct {
}

func (s LvtData) ListByPage(cIds []int32, page *dto.StatisticalPlayerLtvReq) ([]*po.StatisticalPlayerLtv, error) {
	list := make([]*po.StatisticalPlayerLtv, 0)

	adminDb := conn.GetAdminDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}

	err := adminDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LvtData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s LvtData) ListCount(cIds []int32, page *dto.StatisticalPlayerLtvReq) (int64, error) {
	num := int64(0)

	adminDb := conn.GetAdminDB().Model(&po.StatisticalPlayerLtv{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}

	err := adminDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LvtData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
