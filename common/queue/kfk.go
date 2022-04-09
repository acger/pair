package queue

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"os"
	"time"
)

type KPusher struct {
	sarama.SyncProducer
}

func NewKPusher(address []string) *KPusher {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.RequiredAcks = sarama.WaitForLocal
	p, err := sarama.NewSyncProducer(address, config)

	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return nil
	}

	return &KPusher{p}
}

func (p *KPusher) SendWithTopic(topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", message, err)
		logx.Error(message, err)
	} else {
		fmt.Fprintf(os.Stdout, message+"发送成功，partition=%d, offset=%d \n", part, offset)
		logx.Info(message, part, offset)
	}
}


