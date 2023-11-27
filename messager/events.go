package main

import (
	"math/rand"
	"shared/models"
)

func generateRandomReference(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	word := make([]byte, length)
	for i := range word {
		word[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(word)
}

func constructRequested() *models.Event {
	return models.NewEvent("cli_123123", generateRandomReference(8)).AsRequested()
}

func progressAuthorization(ev *models.Event) (*models.Event, bool) {
	event := ev.AsAuthorized()
	return event, false
}

func progressCapture(ev *models.Event) (*models.Event, bool) {
	event := ev.AsCapture()
	return event, false
}

func progressRefund(ev *models.Event) (*models.Event, bool) {
	event := ev.AsRefund()
	return event, false
}

func progressVoid(ev *models.Event) (*models.Event, bool) {
	event := ev.AsVoid()
	return event, true
}

func progressExpiry(ev *models.Event) (*models.Event, bool) {
	event := ev.AsExpiry()
	return event, true
}
