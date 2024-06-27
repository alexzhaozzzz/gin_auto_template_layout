package user

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewLoginData() *LoginData {
	return &LoginData{}
}

type LoginData struct {
}

func (s LoginData) ListByPage(cIds []int32, page *dto.Login) ([]*po.ReportLogin, error) {
	list := make([]*po.ReportLogin, 0)
	err := conn.GetLogDB().Where("channelid IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s LoginData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetLogDB().Model(&po.ReportLogin{}).Where("channelid IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
