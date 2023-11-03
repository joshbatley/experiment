package main

import (
	"time"
)

type Runner struct {
	Store  EventStore
	ticker *time.Ticker
}

func NewRunner(store EventStore, tps int) *Runner {
	return &Runner{
		Store:  store,
		ticker: time.NewTicker(time.Second / time.Duration(tps)),
	}
}

func (e *Runner) StartUp() {
	go func() {
		for {
			select {
			case <-e.ticker.C:
				Generate(e.Store)
				//log.Print(ev)
			}
		}
	}()
	time.Sleep(time.Second)
}
