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

func progressAuthorization(ev *models.Event) *models.Event {
	return ev.AsAuthorized(10, models.ResponseCodeSuccess)
}

func progressCapture(ev *models.Event) *models.Event {
	return ev.AsCapture(0.3, "")
}

func progressRefund(ev *models.Event) *models.Event {
	return ev.AsRefund(0.1, "")
}

func progressVoid(ev *models.Event) *models.Event {
	return ev.AsVoid("2000")
}

func progressExpiry(ev *models.Event) *models.Event {
	return ev.AsExpiry("2000")
}

//func progressSuccessfulResponse(ev *models.Event) models.ResponseCode {
//	return models.SuccessfulResponseCodes[0]
//}
//
//func progressFailureResponse(ev *models.Event) models.ResponseCode {
//	return models.FailureResponseCodes[0]
//}
//
//func progressInfoResponse(ev *models.Event) models.ResponseCode {
//	return models.InformationResponseCode[0]
//}
//
//func progressFraudResponse(ev *models.Event) models.ResponseCode {
//	return models.FraudResponseCode[0]
//}
