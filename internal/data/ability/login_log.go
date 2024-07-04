package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewLoginLogData() *LoginLogData {
	return &LoginLogData{}
}

type LoginLogData struct {
}

func (s LoginLogData) ListByPage(cIds []int32, page *dto.LoginLogReq) ([]*po.ReportLogin, error) {
	list := make([]*po.ReportLogin, 0)
	err := conn.GetLogDB().Where("channelid IN ?", cIds).Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginLogData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s LoginLogData) ListCount(cIds []int32) (int64, error) {
	num := int64(0)
	err := conn.GetLogDB().Model(&po.ReportLogin{}).Where("channelid IN ?", cIds).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("LoginLogData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
