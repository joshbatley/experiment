package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	utils "shared"
	"shared/models"
	"sort"
	"time"
)

type Setting struct {
	Tps int `json:"tps,omitempty"`
}

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	setting, err := utils.ReadConfig[Setting]("./settings.json")
	if err != nil {
		panic(err)
	}
	log.Print(setting.Tps)
	store := NewInMemory()
	GetUnfinishedPayment(store)

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
	log.Print("", v)
	s.RemoveEvent(v.ID)
	v, err := s.GetRandomEvent()
	if err != nil {
		log.Print("Error")
	}
	log.Print("", v)
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
