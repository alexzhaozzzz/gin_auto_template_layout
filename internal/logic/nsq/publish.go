package nsq

import (
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
)

func PublishMsg(topic string, msg []byte) error {
	client, err := conn.GetNSQCli()
	if err != nil {
		logrus.Errorf("PublishMsg GetNSQCli Err: %s", err)
		return err
	}

	// 发布消息
	err = client.Publish(topic, msg)
	if err != nil {
		logrus.Errorf("PublishMsg Publish Err: %s", err)
		return err
	}

	return nil
}
