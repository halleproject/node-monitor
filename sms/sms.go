package sms

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type content struct {
	Name string `json:"name"`
	Detail string `json:"detail"`
}
func SendSMS(node, detail string, phone string) (string, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "", "")
	if err !=nil {
		return "", err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = "HalleToken"
	request.TemplateCode = "SMS_200703322"

	c:=content{node, detail}
	cByte,_:=json.Marshal(c)
	request.TemplateParam = string(cByte)

	//continue
	request.PhoneNumbers = phone
	resp, err := client.SendSms(request)
	if err != nil {
		fmt.Println(phone, "sms failed : ",err.Error())
		return "",err
	}
	return resp.BaseResponse.String(), nil
}
