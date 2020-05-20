package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"red-east/config"
	"red-east/logging"
	"red-east/queue/mconsumer"
	"syscall"
	"time"
)

var mconfig config.Config
var nsqConfig *nsq.Config
var logger logging.NLogger

func main() {
	//sendMessage()
	consumerMessage()
}
func consumerMessage() {
	var merr error
	mconfig, merr = config.InitConfig()
	if merr != nil {
		fmt.Println(merr.Error())
		os.Exit(1)
	}
	var lerr error
	logger, lerr = logging.InitLogger()
	if lerr != nil {
		fmt.Println(lerr.Error())
		os.Exit(1)
	}
	//三秒轮询一遍
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	StartConsumer()
	<-signalChan
}

func StartConsumer() {
	nsqConfig = nsq.NewConfig()
	nsqConfig.LookupdPollInterval = 3 * time.Second
	consumer, err := nsq.NewConsumer("sms", "channel1", nsqConfig)
	if err != nil {
		logger.Error("fail init consumer msg", err.Error())
		os.Exit(1)
	}
	handler := NewHandler("sms")
	consumer.AddHandler(handler)
	if err := consumer.ConnectToNSQLookupd(mconfig.Nsq.LookupdTcpAddress); err != nil {
		logger.Error("fail connect", err.Error())
		os.Exit(1)
	}
}

func NewHandler(topic string) nsq.Handler {
	var handler nsq.Handler
	if topic == "sms" {
		handler = &mconsumer.MsgHandler{}
	}
	return handler
}

//func sendMessage() {
//	_, err := SendVerifyCodeMsg("86", "15818359718", "123456")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//}
