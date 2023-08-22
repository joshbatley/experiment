package main

import (
	"fmt"
	utils "shared"
	"shared/models"
	"sort"
	"time"
)

type Setting struct {
	Tps int `json:"tps,omitempty"`
}

func main() {
	setting, err := utils.ReadConfig[Setting]("./settings.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(setting.Tps)
	store := NewInMemory()
	GetUnfinishedPayment(store)
	//fmt.Printf("%+v\n", event.ID)

}

const (
	ExpireAfter = time.Second * 3000
)

func GenerateEvent() {
	// Random Chance
	// create new event
	// publish
	// grab unfinished event
	//
}

func GetUnfinishedPayment(s EventStore) {
	event := generateAuthorize()
	s.AddUnfinishedEvent(Event{
		ID: "asdasd",
		PastEvents: []models.Payment{
			*event,
		},
	})

	v, _ := s.GetRandomEvent()
	fmt.Print("", v)
	s.RemoveEvent(v.ID)
	v, err := s.GetRandomEvent()
	if err != nil {
		fmt.Print("Error")
	}
	fmt.Print("", v)
	// Get random event from store (not lifo/fifo)
	// Workout out what is next possible actions
	// Generate event
	// Remove old
}

func ShouldExpireEvent(e Event) bool {
	sort.Slice(e.PastEvents, func(i, j int) bool {
		return e.PastEvents[i].Timestamp.Before(e.PastEvents[j].Timestamp)
	})
	newest := e.PastEvents[0].Timestamp
	return newest.After(time.Now().Add(ExpireAfter))
}
