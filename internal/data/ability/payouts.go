package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/dto"
	"gorm.io/gorm"
	"time"
)

func NewPayoutsData() *PayoutsData {
	return &PayoutsData{}
}

type PayoutsData struct {
}

func (s PayoutsData) ListByPage(cIds []int32, page *dto.PayoutsReq) ([]*po.WithdrawRecord, error) {
	list := make([]*po.WithdrawRecord, 0)

	playerDb := conn.GetPlayerDB().Where("n_channel IN ?", cIds)
	if page.OrderStartTime > 0 && page.OrderEndTime > 0 {
		playerDb = playerDb.Where("n_createtime BETWEEN ? AND ?", time.Unix(page.OrderStartTime, 0), time.Unix(page.OrderEndTime, 0))
	}
	if page.LoanStartTime > 0 && page.LoanEndTime > 0 {
		playerDb = playerDb.Where("n_paytime BETWEEN ? AND ?", time.Unix(page.LoanStartTime, 0), time.Unix(page.LoanEndTime, 0))
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}
	if page.State > 0 {
		playerDb = playerDb.Where("n_state = ?", page.State)
	}

	err := playerDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListByPage Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PayoutsData) ListCount(cIds []int32, page *dto.PayoutsReq) (int64, error) {
	num := int64(0)
	playerDb := conn.GetPlayerDB().Model(&po.WithdrawRecord{}).Where("n_channel IN ?", cIds)
	if page.OrderStartTime > 0 && page.OrderEndTime > 0 {
		playerDb = playerDb.Where("n_createtime BETWEEN ? AND ?", time.Unix(page.OrderStartTime, 0), time.Unix(page.OrderEndTime, 0))
	}
	if page.LoanStartTime > 0 && page.LoanEndTime > 0 {
		playerDb = playerDb.Where("n_paytime BETWEEN ? AND ?", time.Unix(page.LoanStartTime, 0), time.Unix(page.LoanEndTime, 0))
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}
	if page.State > 0 {
		playerDb = playerDb.Where("n_state = ?", page.State)
	}

	err := playerDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PayoutsData) ListByStatePage(cIds []int32, state int32, page *dto.PayoutsAuditReq) ([]*po.WithdrawRecord, error) {
	list := make([]*po.WithdrawRecord, 0)
	playerDb := conn.GetPlayerDB().Where("n_channel IN ? AND n_state = ? ", cIds, state)
	if page.AmountStart > 0 && page.AmountEnd > 0 {
		playerDb = playerDb.Where("n_amount BETWEEN ? AND ?", page.AmountStart, page.AmountEnd)
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}

	err := playerDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListByStatePage Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PayoutsData) ListStateCount(cIds []int32, state int32, page *dto.PayoutsAuditReq) (int64, error) {
	num := int64(0)
	playerDb := conn.GetPlayerDB().Model(&po.WithdrawRecord{}).Where("n_channel IN ? AND n_state = ? ", cIds, state)
	if page.AmountStart > 0 && page.AmountEnd > 0 {
		playerDb = playerDb.Where("n_amount BETWEEN ? AND ?", page.AmountStart, page.AmountEnd)
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}

	err := playerDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListStateCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s PayoutsData) ListByPlayerPage(cIds []int32, playerId int64, page *dto.UserListReq) ([]*po.WithdrawRecord, error) {
	list := make([]*po.WithdrawRecord, 0)
	playerDb := conn.GetPlayerDB().Where("n_channel IN ? AND n_playerid = ? ", cIds, playerId)

	//TODO:需要增加时间查询
	err := playerDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListByPlayerPage Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s PayoutsData) ListPlayerCount(cIds []int32, playerId int64) (int64, error) {
	num := int64(0)
	playerDb := conn.GetPlayerDB().Model(&po.WithdrawRecord{}).Where("n_channel IN ? AND n_playerid = ? ", cIds, playerId)

	err := playerDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("PayoutsData ListPlayerCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}
