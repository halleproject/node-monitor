package sms

import (
	"fmt"
	"testing"
	)

func TestSendSMS(t *testing.T) {
	fmt.Println("测试短信发送")

	SendSMS("测试节点 name","测试节点 detail", "13021248048")
}
