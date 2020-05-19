package producer

import (
	"encoding/json"
	"fmt"
	. "red-east/queue/queue"
	"time"
)

type Job struct {
	Time int64       `json:"time"`
	Args interface{} `json:"args"`
}

type sendMsgArgs struct {
	CountryCode string `json:"country_code"`
	Mobile      string `json:"mobile"`
	Msg         string `json:"msg"`
	Code        string `json:"code"`
}

func SendVerifyCodeMsg(countryCode, mobile, code string) (bool, error) {
	msg := fmt.Sprintf("您本次密码设置操作的验证码为：%s，有效期30分钟。", code)
	msgArgs := sendMsgArgs{
		CountryCode: countryCode,
		Mobile:      mobile,
		Msg:         msg,
		Code:        code,
	}
	job := Job{
		Time: time.Now().Unix(),
		Args: msgArgs,
	}
	data, err := json.Marshal(job)
	if err != nil {
		return false, err
	}
	if err := PushJob("sms", data); err != nil {
		fmt.Println(err.Error())
		return false, nil
	}
	return true, nil
}
