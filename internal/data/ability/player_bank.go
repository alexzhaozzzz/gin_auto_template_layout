package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewPlayerBankData() *PlayerBankData {
	return &PlayerBankData{}
}

type PlayerBankData struct {
}

//func (s PlayerBankData) ListByPage(page *dto.PlayerBank) ([]*po.PlayerBank, error) {
//	list := make([]*po.PlayerBank, 0)
//	err := conn.GetPlayerDB().Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		logrus.Errorf("PlayerBankData List Err: %s", err.Error())
//		return nil, err
//	}
//	return list, nil
//}

func (s PlayerBankData) ListCount() (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.PlayerBank{}).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerBankData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PlayerBankData) InfoById(id uint64) (*po.PlayerBank, error) {
	info := &po.PlayerBank{}
	err := conn.GetPlayerDB().First(info, "n_playerid = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerBankData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
