package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
)

func AddTestConsumer(topic, channel string) error {
	client, err := conn.GetNSQCli()
	if err != nil {
		logrus.Errorf("AddTestConsumer GetNSQCli Err: %s", err)
		return err
	}

	// 添加消费者示例（这里简化处理，实际应根据需求实现handler）
	handler := nsq.HandlerFunc(func(message *nsq.Message) error {
		logrus.Errorf("AddTestConsumer : %s", string(message.Body))
		return nil
	})
	err = client.AddConsumer(topic, channel, handler)
	if err != nil {
		logrus.Errorf("AddTestConsumer Failed to add consumer Err: %s", err)
		return err
	}

	return nil
}
