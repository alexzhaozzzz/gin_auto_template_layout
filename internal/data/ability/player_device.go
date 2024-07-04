package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewPlayerDeviceData() *PlayerDeviceData {
	return &PlayerDeviceData{}
}

type PlayerDeviceData struct {
}

//func (s PlayerDeviceData) ListByPage(page *dto.PlayerDevice) ([]*po.PlayerDevice, error) {
//	list := make([]*po.PlayerDevice, 0)
//	err := conn.GetLogDB().Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		logrus.Errorf("PlayerDeviceData List Err: %s", err.Error())
//		return nil, err
//	}
//	return list, nil
//}

func (s PlayerDeviceData) ListCount() (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.PlayerDevice{}).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerDeviceData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PlayerDeviceData) InfoById(id uint64) (*po.PlayerDevice, error) {
	info := &po.PlayerDevice{}
	err := conn.GetPlayerDB().First(info, "n_player_id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerDeviceData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
