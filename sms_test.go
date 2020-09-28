package sms

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	fmt.Println("测试短信发送")

	res, err := SendSMS("测试节点 name", "123456667",
		"15611084264")
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
