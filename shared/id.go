package utils

import (
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
)

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
