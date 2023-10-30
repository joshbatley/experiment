package utils

import (
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"os"
)

func NewEventId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func NewActionId(eventId string) string {
	return "act_" + base64.StdEncoding.EncodeToString([]byte(eventId))
}

func NewPaymentId() string {
	return "pay_" + base64.StdEncoding.EncodeToString([]byte(NewEventId()))
}

func FindIndex[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func ReadConfig[T any]() (T, error) {
	return readConfig[T]("./settings.json")
}

func readConfig[T any](path string) (T, error) {
	var val T
	data, err := os.ReadFile(path)
	if err != nil {
		return val, err
	}

	if err := json.Unmarshal(data, &val); err != nil {
		return val, err
	}
	return val, nil
}
