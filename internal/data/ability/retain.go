package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewRetainData() *RetainData {
	return &RetainData{}
}

type RetainData struct {
}

func (s RetainData) ListByPage(cIds []int32, page *dto.RetainReq) ([]*po.Retain, error) {
	list := make([]*po.Retain, 0)

	adminDb := conn.GetAdminDB().Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.RType > 0 {
		adminDb.Where("type = ?", page.RType)
	}

	err := adminDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RetainData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RetainData) ListCount(cIds []int32, page *dto.RetainReq) (int64, error) {
	num := int64(0)

	adminDb := conn.GetAdminDB().Model(&po.Retain{}).Where("channel_id IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		adminDb.Where("date BETWEEN ? AND ?", page.StartTime, page.EndTime)
	}
	if page.RType > 0 {
		adminDb.Where("type = ?", page.RType)
	}

	err := adminDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RetainData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
