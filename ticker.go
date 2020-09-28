package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tendermint/tendermint/libs/log"
)

var (
	tickTockBufferSize = 100000
)

type AlarmType uint8

func (a AlarmType) String() string {
	switch a {
	case 0:
		return "query"
	case 1:
		return "Block"
	case 2:
		return "Val"
	default:
		return "just support 0,1,2"
	}
}

const (
	Query AlarmType = iota
	BlockHeight
	Validator
)

type AlarmInfo struct {
	AlarmType AlarmType
	Detail    string
	Time      time.Time
}

func (ai *AlarmInfo) String() string {
	return fmt.Sprintf("alarm type: %s  detail: %s  time: %s ", ai.AlarmType, ai.Detail, ai.Time.Format("2006-01-02 15:04:05"))
}

type TimerTicker struct {
	ticker *time.Ticker
	Data   chan AlarmInfo
	Logger log.Logger
	Cli    *Client
}

func NewTimerticker(cli *Client) *TimerTicker {
	tt := &TimerTicker{
		Cli:    cli,
		ticker: time.NewTicker(time.Second),
		Data:   make(chan AlarmInfo, tickTockBufferSize),
		Logger: log.NewTMLogger(os.Stdout),
	}
	//tt.stopTimer() // don't want to fire until the first scheduled timeout
	return tt
}

func (t *TimerTicker) OnStart() error {
	go t.timeoutRoutine()
	return nil
}

// stop the timer and drain if necessary
func (t *TimerTicker) stopTimer() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
}

func (t *TimerTicker) timeoutRoutine() {
	t.Logger.Debug("Starting timeout routine")
	var ti time.Time
	var count uint64

	for {
		select {
		case newti := <-t.ticker.C:
			t.Logger.Debug("Received tick", "old_ti", ti, "new_ti", newti)
			ti = newti

			l := len(t.Data)

			if l == 0 {
				count = 0
				t.ticker = time.NewTicker(time.Second)
				continue
			}

			datas := make([]AlarmInfo, 0)

			t.Logger.Debug("datas len", "len", l)

			for i := 0; i < l; i++ {
				data, ok := <-t.Data
				if !ok {
					return
				}
				datas = append(datas, data)
			}

			count++

			switch count {
			case 1:
				t.ticker = time.NewTicker(time.Minute)
			case 3:
				t.ticker = time.NewTicker(10 * time.Minute)
			case 6:
				t.ticker = time.NewTicker(time.Hour)
			}

			t.Logger.Debug("send msg.", "times", count)
			res, err := t.Cli.SendAlarm(datas)
			if err != nil {
				t.Logger.Error(err.Error())
			}
			t.Logger.Debug(res, "times", count)
		}
	}
}
