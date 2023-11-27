package utils

import (
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

func FindIndex[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Find[T comparable](slice []T, comparable func(T) bool) (T, error) {
	for _, v := range slice {
		if comparable(v) {
			return v, nil
		}
	}
	var empty T
	return empty, errors.New("not found")
}

func GetRandomItem[T any](slice []T) T {
	randomIdx := rand.Intn(len(slice))
	return slice[randomIdx]
}

func GenerateRandomNumber() float64 {
	randomNumber := rand.Float64() * 100000
	return randomNumber
}

func GenerateRandomNumberBetween(max float64) float64 {
	randomNumber := rand.Float64() * max
	if randomNumber == 0 {
		return 0.1
	}
	return randomNumber
}

func NewEventId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func NewActionId(eventId string) string {
	return "act_" + strings.ToLower(base64.StdEncoding.EncodeToString([]byte(eventId)))
}

func NewPaymentId() string {
	return "pay_" + strings.ToLower(base64.StdEncoding.EncodeToString([]byte(NewEventId())))
}

func GenerateRandomReference(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	word := make([]byte, length)
	for i := range word {
		word[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(word)
}
