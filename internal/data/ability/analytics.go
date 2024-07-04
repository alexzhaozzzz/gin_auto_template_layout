package ability

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/po"
	"gorm.io/gorm"
)

func NewAnalyticsData() *AnalyticsData {
	return &AnalyticsData{}
}

type AnalyticsData struct {
}

func (s AnalyticsData) InfoByChannelId(id int64) (*po.GoogleAnalytics, error) {
	info := &po.GoogleAnalytics{}
	err := conn.GetAdminDB().First(info, "channel_id = ?", id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.Errorf("ChannelData InfoByChannelId Err: %s", err.Error())
		return nil, err
	}
	return info, nil
}
