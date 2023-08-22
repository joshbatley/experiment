package models

type Action string

const (
	ActionAuthorize Action = "Authorize"
	ActionCapture   Action = "Capture"
	ActionRefund    Action = "Refund"
	ActionExpiry    Action = "Expiry"
	ActionVoid      Action = "Void"
)
