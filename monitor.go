package main

import (
	"fmt"
	"time"
)

func (cli *Client) Monitor() {

	lastHeight := int64(0)
	lastTime := time.Now()

	apiErrTry := 0
	for {
		if apiErrTry > 5 {
			fmt.Errorf("API failed 3 times, send alarm !!!")
			//processAlarm(time.Now().Unix())
			apiErrTry = 0
		}

		height, err := cli.LatestBlockHeight()
		if err != nil {
			fmt.Errorf("failed to query block height")
			apiErrTry++
			time.Sleep(1 * time.Second)
			continue
		}
		if lastHeight == height {
			alarmTime := time.Now()
			elapse := alarmTime.Sub(lastTime).Seconds()
			if elapse < 120 {
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Errorf("Height %v is not right, take time: %v seconds , send alarm !!!", height, elapse)
			//processAlarm(alarmTime.Unix())
			continue
		}
		lastHeight = height
		lastTime = time.Now()

		vals, err := cli.GetValidators()
		if err != nil {
			apiErrTry++
			time.Sleep(1 * time.Second)
			continue
		}

		for _, v := range vals {
			if v.Status != 3 {
				fmt.Errorf("Validators status !=3  block: %v, send alarm !!!", height)
				//processAlarm(time.Now().Unix())
			}
		}
		apiErrTry = 0
	}
}
