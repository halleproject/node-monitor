package main

import (
	"time"

	"github.com/tendermint/tendermint/libs/log"
)

var (
	tickTockBufferSize = 10
)

type TimerTicker struct {
	timer  *time.Timer
	ticker *time.Ticker
	Data   chan []byte
	Logger log.Logger
	// tickChan chan timeoutInfo // for scheduling timeouts
	// tockChan chan timeoutInfo // for notifying about them
}

func NewTimerticker() *TimerTicker {
	tt := &TimerTicker{
		timer: time.NewTimer(0),
		//ticker: time.NewTicker(0),
		Data: make(chan []byte, tickTockBufferSize),
	}
	tt.stopTimer() // don't want to fire until the first scheduled timeout
	return tt
}

func (t *TimerTicker) OnStart() error {
	go t.timeoutRoutine()
	return nil
}

func (t *TimerTicker) ScheduleTimeout(d time.Duration) {
	t.timer.Reset(d)
}

// stop the timer and drain if necessary
func (t *TimerTicker) stopTimer() {
	// Stop() returns false if it was already fired or was stopped
	if !t.timer.Stop() {
		select {
		case <-t.timer.C:
		default:
			t.Logger.Debug("Timer already stopped")
		}
	}
	t.ticker.Stop()
}

func (t *TimerTicker) timeoutRoutine() {
	t.Logger.Debug("Starting timeout routine")
	var ti time.Time
	var datas [][]byte
	var count uint64

	for {
		select {
		case data, ok := <-t.Data:
			if !ok {
				return
			}
			if count == 0 {
				count = 1
				// send msg
			} else {
				datas = append(datas, data)
			}
		case newti := <-t.ticker.C:
			t.Logger.Debug("Received tick", "old_ti", ti, "new_ti", newti)
			ti = newti

			if len(datas) > 0 {
				//send msg
				count++
			} else {
				count = 0
			}
		case <-t.timer.C:

			if count == 1 {
				t.ticker = time.NewTicker(time.Minute)
				t.ScheduleTimeout(4 * time.Minute)
			}

			if count > 3 && count < 5 {
				t.ticker = time.NewTicker(10 * time.Minute)
				t.ScheduleTimeout(time.Hour)
			}

			if count > 8 {
				t.ticker = time.NewTicker(time.Hour)
			}
		}
	}
}
