package system

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewMerchantChannelBindData() *MerchantChannelBindData {
	return &MerchantChannelBindData{}
}

type MerchantChannelBindData struct {
}

func (s MerchantChannelBindData) ListByMerchantId(mId uint64) ([]*po.MerchantChannelBind, error) {
	list := make([]*po.MerchantChannelBind, 0)
	err := conn.GetMerchantDB().Where("merchant_id = ?", mId).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("MerchantChannelBindData List Err: %s", err.Error())
		return nil, err
	}
	return list, nil
}
