package main

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func (c *Client) SendAlarm(datas []AlarmInfo) (string, error) {
	c.Logger.Debug("param", "datas len", len(datas))

	if len(datas) == 1 {

		c.Logger.Debug(datas[0].String())
		return SendSMS(c.Moniker, datas[0].AlarmType.String(), "15611084264,18001052350")
	}

	var qC, bC, vC uint
	for _, v := range datas {
		switch v.AlarmType {
		case Query:
			qC++
		case BlockHeight:
			bC++
		case Validator:
			vC++
		}
	}

	c.Logger.Debug(fmt.Sprintf("%v%v", c.Moniker, fmt.Sprintf("query: %d block: %d val: %d", qC, bC, vC)))

	//return timeStr, nil
	return SendSMS(c.Moniker, fmt.Sprintf("q:%db:%dv:%d", qC, bC, vC), "15611084264,18001052350")
}

type content struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func SendSMS(node, detail string, phone string) (string, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "", "")
	if err != nil {
		return "", err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = "HalleToken"
	request.TemplateCode = "SMS_200703322"

	cnt := content{node, detail}
	cByte, _ := json.Marshal(cnt)
	request.TemplateParam = string(cByte)

	//continue
	request.PhoneNumbers = phone
	resp, err := client.SendSms(request)
	if err != nil {
		fmt.Println(phone, "sms failed : ", err.Error())
		return "", err
	}
	return resp.BaseResponse.String(), nil
}
