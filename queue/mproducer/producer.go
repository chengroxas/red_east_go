package mproducer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"red-east/config"
)

func InitProducer() *nsq.Producer {
	config, _ := config.InitConfig()
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(config.Nsq.TcpAddress, nsqConfig)
	if err != nil {
		fmt.Println("init producer fail:", err.Error())
		return nil
	}
	return producer
}

func PushJob(topic string, data []byte) error {
	producer := InitProducer()
	err := producer.Publish(topic, data)
	if err != nil {
		return err
	}
	return nil
}
