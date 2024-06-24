package nsq

import (
	"github.com/sirupsen/logrus"
)

func PublishMsg(topic string, msg []byte) error {
	client, err := data.GetNSQCli()
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
