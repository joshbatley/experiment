package main

import (
	"shared"
	"shared/models"
)

func generateAuthorize() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsAuthorized(10, models.ResponseCodeSuccess)
}

func generateCapture() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func generateRefund() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func generateVoid() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func generateSuccessfulResponse() models.ResponseCode {
	return models.SuccessfulResponseCodes[0]
}

func generateFailureResponse() models.ResponseCode {
	return models.FailureResponseCodes[0]
}

func generateInfoResponse() models.ResponseCode {
	return models.InformationResponseCode[0]
}

func generateFraudResponse() models.ResponseCode {
	return models.FraudResponseCode[0]
}
