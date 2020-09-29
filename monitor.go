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
			cli.Logger.Debug("query api failed, send alarm !!!", "try times", apiErrTry)

			alarm := AlarmInfo{
				AlarmType: Query,
				Detail:    fmt.Sprintf("query failed, try times: %v", apiErrTry),
				Time:      time.Now(),
			}
			cli.timerTicker.Data <- alarm
			apiErrTry = 0
		}

		height, err := cli.LatestBlockHeight()
		if err != nil {
			cli.Logger.Debug("faild to query block height")
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
			cli.Logger.Debug(fmt.Sprintf("Block height %v is not increased, take time: %v seconds.", height, elapse))
			alarm := AlarmInfo{
				AlarmType: BlockHeight,
				Detail:    fmt.Sprintf("Block %v doesnot increased, take time: %v seconds.", height, elapse),
				Time:      time.Now(),
			}
			cli.timerTicker.Data <- alarm
			time.Sleep(5 * time.Second)

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
				cli.Logger.Debug(fmt.Sprintf("Validators status !=3  block: %v", height))

				alarm := AlarmInfo{
					AlarmType: Validator,
					Detail:    fmt.Sprintf("ValSts %v blk %v", v.Status, height),
					Time:      time.Now(),
				}
				cli.timerTicker.Data <- alarm

			}
		}
		apiErrTry = 0
	}
}
