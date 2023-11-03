package main

import (
	"github.com/rs/zerolog/log"
	"shared/models"
	"sort"
	"time"
)

const (
	ExpireAfter = time.Second * 3000
)

var count int = 0

func Generate(store EventStore) string {
	count++
	log.Print("Record fired", count)
	return ""
}

func GenerateEvent() {
	// Random Chance
	// create new event
	// publish
	// grab unfinished event
	//
}

func GetUnfinishedPayment(s EventStore) {
	event := constructRequested()
	s.AddUnfinishedEvent(Record{
		ID: "asdasd",
		PastEvents: []models.Event{
			*event,
		},
	})

	v, _ := s.GetRandomEvent()
	//log.Print("", v)
	s.RemoveEvent(v.ID)
	v, err := s.GetRandomEvent()
	if err != nil {
		//log.Print("Error")
	}
	//log.Print("", v)
	// Get random event from store (not lifo/fifo)
	// Workout out what is next possible actions
	// Generate event
	// Remove old
}

func ShouldExpireEvent(e Record) bool {
	sort.Slice(e.PastEvents, func(i, j int) bool {
		return e.PastEvents[i].Timestamp.Before(e.PastEvents[j].Timestamp)
	})
	newest := e.PastEvents[0].Timestamp
	return newest.After(time.Now().Add(ExpireAfter))
}
