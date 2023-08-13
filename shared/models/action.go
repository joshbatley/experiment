package models

type Action string

const (
	ActionAuthorize Action = "Authorize"
	ActionCapture   Action = "Capture"
	ActionRefund    Action = "Refund"
	ActionVoid      Action = "Void"
)
