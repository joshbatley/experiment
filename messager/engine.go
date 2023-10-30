package main

import (
	"time"
)

type Engine struct {
	Store  EventStore
	ticker *time.Ticker
}

func NewEngine(store EventStore, tps int) *Engine {
	return &Engine{
		Store:  store,
		ticker: time.NewTicker(time.Second / time.Duration(tps)),
	}
}

func (e *Engine) StartUp() {
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
