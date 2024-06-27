package statistics

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewLoginRetainData() *LoginRetainData {
	return &LoginRetainData{}
}

type LoginRetainData struct {
}

func (s LoginRetainData) ListByPage(cIds []int32, page *dto.Retain) ([]*po.Retain, error) {
	list := make([]*po.Retain, 0)
	err := conn.GetAdminDB().Where("channel_id IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginRetainData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s LoginRetainData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetAdminDB().Model(&po.Retain{}).Where("channel_id IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginRetainData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
