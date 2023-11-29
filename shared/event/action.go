package event

type Action string

const (
	ActionRequested Action = "Requested"
	ActionAuthorize Action = "Authorize"
	ActionCapture   Action = "Capture"
	ActionRefund    Action = "Refund"
	ActionExpiry    Action = "Expiry"
	ActionVoid      Action = "Void"
)
