package common

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/system"
)

func GetChannelIdsByMerchantId(mIds uint64) ([]int32, error) {
	ids := make([]int32, 0)

	d := system.NewMerchantChannelBindData()
	list, err := d.ListByMerchantId(mIds)
	if err != nil {
		logrus.Errorf("GetChannelIdsByMerchantId Err: %s", err.Error())
		return nil, err
	}
	if len(list) > 0 {
		for _, v := range list {
			ids = append(ids, v.ChannelId)
		}
	}

	return ids, nil
}
