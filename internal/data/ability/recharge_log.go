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

func NewRechargeLogData() *RechargeLogData {
	return &RechargeLogData{}
}

type RechargeLogData struct {
}

func (s RechargeLogData) ListByPage(cIds []int32, page *dto.RechargeLogReq) ([]*po.RechargeRecord, error) {
	list := make([]*po.RechargeRecord, 0)
	playerDb := conn.GetPlayerDB().Where("n_channel IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		playerDb = playerDb.Where("n_createtime BETWEEN ? AND ?", time.Unix(page.StartTime, 0), time.Unix(page.EndTime, 0))
	}
	if page.FinishStartTime > 0 && page.FinishEndTime > 0 {
		playerDb = playerDb.Where("n_finishtime BETWEEN ? AND ?", time.Unix(page.FinishStartTime, 0), time.Unix(page.FinishEndTime, 0))
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}
	if page.Amount != "" {
		playerDb = playerDb.Where("n_amount >= ?", page.Amount)
	}
	if page.Platform != "" {
		playerDb = playerDb.Where("n_platform = ?", page.Platform)
	}
	if page.State > 0 {
		playerDb = playerDb.Where("n_state = ?", page.State)
	}

	err := playerDb.Offset(page.GetOffset()).Limit(page.GetPageSize()).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RechargeLogData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

func (s RechargeLogData) ListCount(cIds []int32, page *dto.RechargeLogReq) (int64, error) {
	num := int64(0)
	playerDb := conn.GetPlayerDB().Model(&po.RechargeRecord{}).Where("n_channel IN ?", cIds)
	if page.StartTime > 0 && page.EndTime > 0 {
		playerDb = playerDb.Where("n_createtime BETWEEN ? AND ?", time.Unix(page.StartTime, 0), time.Unix(page.EndTime, 0))
	}
	if page.FinishStartTime > 0 && page.FinishEndTime > 0 {
		playerDb = playerDb.Where("n_finishtime BETWEEN ? AND ?", time.Unix(page.FinishStartTime, 0), time.Unix(page.FinishEndTime, 0))
	}
	if page.PlayerId > 0 {
		playerDb = playerDb.Where("n_playerid = ?", page.PlayerId)
	}
	if page.OrderId != "" {
		playerDb = playerDb.Where("n_orderid = ?", page.OrderId)
	}
	if page.Amount != "" {
		playerDb = playerDb.Where("n_amount >= ?", page.Amount)
	}
	if page.Platform != "" {
		playerDb = playerDb.Where("n_platform = ?", page.Platform)
	}
	if page.State > 0 {
		playerDb = playerDb.Where("n_state = ?", page.State)
	}

	err := playerDb.Count(&num).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RechargeLogData ListCount Err: %s", err.Error())
		return num, err
	}
	return num, nil
}

func (s RechargeLogData) InfoByOrderId(id string) (*po.RechargeRecord, error) {
	info := &po.RechargeRecord{}
	err := conn.GetPlayerDB().First(info, "n_orderid = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("RechargeLogData InfoByOrderId Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
