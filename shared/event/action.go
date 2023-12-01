package event

type Action string

const (
	ActionRequest   Action = "Request"
	ActionAuthorize Action = "Authorize"
	ActionCapture   Action = "Capture"
	ActionRefund    Action = "Refund"
	ActionExpiry    Action = "Expiry"
	ActionVoid      Action = "Void"
)
