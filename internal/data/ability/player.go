package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewPlayerData() *PlayerData {
	return &PlayerData{}
}

type PlayerData struct {
}

//func (s PlayerData) ListByPage(page *dto.Player) ([]*po.Player, error) {
//	list := make([]*po.Player, 0)
//	err := conn.GetPlayerDB().Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		logrus.Errorf("PlayerData List Err: %s", err.Error())
//		return nil, err
//	}
//	return list, nil
//}

func (s PlayerData) ListCount() (int64, error) {
	num := int64(0)
	err := conn.GetPlayerDB().Model(&po.Player{}).Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerData List Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PlayerData) InfoById(id int64) (*po.Player, error) {
	info := &po.Player{}
	err := conn.GetPlayerDB().First(info, "n_player_id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
