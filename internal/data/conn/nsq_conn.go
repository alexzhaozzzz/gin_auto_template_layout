package conn

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

// 消息超时
var pubMsgTimeout = 3 * time.Minute

var NSQClient *NSQ

func GetNSQCli() (*NSQ, error) {
	var nsqCli *NSQ
	if NSQClient == nil {
		nsqc, err := newNSQClient()
		if err != nil {
			logrus.Errorf("GetNSQCli newNSQClient Err: %s", err.Error())
			return nil, err
		}

		nsqCli = nsqc
	} else {
		nsqCli = NSQClient
	}

	return nsqCli, nil
}

// NSQ 封装了NSQ的生产和消费功能
type NSQ struct {
	producer  *nsq.Producer
	consumers []*nsq.Consumer
}

// newNSQClient 初始化NSQClient
func newNSQClient() (*NSQ, error) {
	addr := viper.GetString("nsq.addr")
	//lookUpAddr := viper.GetStringSlice("nsq.lookUpdAddr")

	config := nsq.NewConfig()
	config.MsgTimeout = pubMsgTimeout

	// 初始化生产者
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}

	producer.SetLogger(nil, nsq.LogLevelWarning)
	if err := producer.Ping(); err != nil {
		panic("conn to nsq err")
	}

	client := &NSQ{
		producer: producer,
	}

	// 如果需要，可以在这里初始化消费者并添加到consumers列表中
	// 这里省略了具体的消费者初始化逻辑，因为具体实现依赖于应用需求

	NSQClient = client

	return client, nil
}

// Publish 发布消息到指定的主题
func (nc *NSQ) Publish(topic string, body []byte) error {
	return nc.producer.Publish(topic, body)
}

// Close 关闭所有连接
func (nc *NSQ) Close() error {
	var err error
	for _, consumer := range nc.consumers {
		consumer.Stop()
	}

	nc.producer.Stop()
	return err
}

// AddConsumer 添加消费者
// 这个方法应该根据实际需求实现消费者逻辑，这里仅作为示例
func (nc *NSQ) AddConsumer(topic, channel string, handler nsq.Handler) error {
	consumer, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return err
	}
	consumer.AddHandler(handler)

	err = consumer.ConnectToNSQD(nc.producer.String())
	if err != nil {
		return err
	}
	nc.consumers = append(nc.consumers, consumer)
	return nil
}
