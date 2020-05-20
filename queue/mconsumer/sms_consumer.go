package mconsumer

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"red-east/queue/mproducer"
	"time"
)

type MsgHandler struct {
}

//当message.Attempts > r.config.MaxAttempts时，就会调用改方法进行处理
func (handler *MsgHandler) LogFailedMessage(message *nsq.Message) {
	fmt.Println("不要再重试了，记录到数据库中")
}

//根据情况调用message.Requeue将message重新放回队列中，
//重新尝试r.config.MaxAttempts次后就会调用LogFailedMessage处理
func (handler *MsgHandler) HandleMessage(message *nsq.Message) error {
	job := mproducer.Job{}
	fmt.Println(message)
	if err := json.Unmarshal(message.Body, &job); err != nil {
		//数据问题，属于代码逻辑错误，直接抛弃不要
		fmt.Println("msg接受body数据问题:", err.Error(), " body:", string(message.Body))
		return err
	}
	//time.Sleep(10 * time.Second) //耗时的消耗
	fmt.Println("处理完成了")
	fmt.Println(job)
	if err := sendMsg(); err != nil {
		//重新发送，根据短信方返回内容决定是否重发发送
		message.Requeue(2 * time.Second)
	}
	return nil
}

func sendMsg() error {
	return nil
}
