package models

type Status string

const (
	StatusPending           Status = "Pending"
	StatusCompleted         Status = "Completed"
	StatusCancelled         Status = "Cancelled"
	StatusFailed            Status = "Failed"
	StatusCaptured          Status = "Captured"
	StatusPartiallyCaptured Status = "PartiallyCaptured"
	StatusRefunded          Status = "Refunded"
	StatusPartiallyRefunded Status = "PartiallyRefunded"
	StatusAuthorized        Status = "Authorized"
)
