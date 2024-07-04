package ability

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
)

func NewGameLogData() *GameLogData {
	return &GameLogData{}
}

type GameLogData struct {
}

func (s GameLogData) GetTableName(day string) string {
	cr := po.GameLog{}
	table := fmt.Sprintf("%s_%s", cr.TableName(), day)
	return table
}

func (s GameLogData) ListByPage(day string, cIds []int32, page *dto.GameLogReq) ([]*po.GameLog, error) {
	list := make([]*po.GameLog, 0)
	logDb := conn.GetLogDB().Table(s.GetTableName(day)).Where("channel_id IN ?", cIds)
	if page.PlayerId > 0 {
		logDb = logDb.Where("player_id = ?", page.PlayerId)
	}
	if page.GameId > 0 {
		logDb = logDb.Where("game_id = ?", page.GameId)
	}
	if page.Bet > 0 {
		logDb = logDb.Where("bet = ?", page.Bet)
	}
	if page.Reward > 0 {
		logDb = logDb.Where("reward = ?", page.Reward)
	}
	if page.Win > 0 {
		logDb = logDb.Where("win = ?", page.Win)
	}
	if page.LogType > 0 {
		logDb = logDb.Where("log_type = ?", page.LogType)
	}

	err := logDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("GameLogData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s GameLogData) ListCount(day string, cIds []int32, page *dto.GameLogReq) (int64, error) {
	num := int64(0)
	logDb := conn.GetLogDB().Table(s.GetTableName(day)).Model(&po.GameLog{}).Where("channel_id IN ?", cIds)
	if page.PlayerId > 0 {
		logDb = logDb.Where("player_id = ?", page.PlayerId)
	}
	if page.GameId > 0 {
		logDb = logDb.Where("game_id = ?", page.GameId)
	}
	if page.Bet > 0 {
		logDb = logDb.Where("bet = ?", page.Bet)
	}
	if page.Reward > 0 {
		logDb = logDb.Where("reward = ?", page.Reward)
	}
	if page.Win > 0 {
		logDb = logDb.Where("win = ?", page.Win)
	}
	if page.LogType > 0 {
		logDb = logDb.Where("log_type = ?", page.LogType)
	}

	err := logDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("GameLogData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
