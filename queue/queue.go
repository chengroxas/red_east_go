package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"red-east/config"
	"red-east/logging"
	"red-east/queue/consumer"
	. "red-east/queue/producer"
	"syscall"
	"time"
)

func main() {
	mconfig, merr := config.InitConfig()
	if merr != nil {
		fmt.Println(merr.Error())
		os.Exit(1)
	}
	logger, lerr := logging.InitLogger()
	if lerr != nil {
		fmt.Println(lerr.Error())
		os.Exit(1)
	}
	nsqConfig := nsq.NewConfig()
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			consume, err := nsq.NewConsumer("sms", "channel1", nsqConfig)
			if err != nil {
				logger.Error("fail init consumer msg", err.Error())
				os.Exit(1)
			}
			handler := consumer.MsgHandler{}
			consume.AddHandler(&handler)
			if err := consume.ConnectToNSQLookupd(mconfig.Nsq.LookupdTcpAddress); err != nil {
				logger.Error("fail connect", err.Error())
				os.Exit(1)
			}
		}
	}()
	<-signalChan
}

func sendMessage() {
	_, err := SendVerifyCodeMsg("86", "15818359718", "123456")
	if err != nil {
		fmt.Println(err.Error())
	}

}
