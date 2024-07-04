package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
)

func NewPlayerProfileData() *PlayerProfileData {
	return &PlayerProfileData{}
}

type PlayerProfileData struct{}

func (s PlayerProfileData) InfoById(id uint64) (*po.PlayerProfile, error) {
	info := &po.PlayerProfile{}
	err := conn.GetPlayerDB().First(info, "n_playerid = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PlayerProfileData InfoById Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
