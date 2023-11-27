package models

type Status string

const (
	StatusPending           Status = "Pending"
	StatusAuthorized        Status = "Authorized"
	StatusCancelled         Status = "Cancelled"
	StatusFailed            Status = "Failed"
	StatusCaptured          Status = "Captured"
	StatusPartiallyCaptured Status = "PartiallyCaptured"
	StatusRefunded          Status = "Refunded"
	StatusPartiallyRefunded Status = "PartiallyRefunded"
)
