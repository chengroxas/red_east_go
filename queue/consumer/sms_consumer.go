package consumer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type MsgHandler struct {
}

func (handler *MsgHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
