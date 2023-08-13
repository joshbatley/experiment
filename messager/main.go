package main

import (
	"fmt"
	"shared/models"
)

func main() {

	store := NewInMemory()
	GetUnfinishedPayment(store)
	//fmt.Printf("%+v\n", event.ID)
}

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
