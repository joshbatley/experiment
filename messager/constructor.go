package main

import (
	"shared"
	"shared/models"
)

func constructAuthorization() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	).AsAuthorized(10, models.ResponseCodeSuccess)
}

func constructCapture() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func constructRefund() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func constructVoid() *models.Payment {
	return models.NewPayment(
		utils.NewEventId(), utils.NewPaymentId(), "", "",
		"", models.CurrencyAUD, models.PaymentMethodApplePay,
	)
}

func constructSuccessfulResponse() models.ResponseCode {
	return models.SuccessfulResponseCodes[0]
}

func constructFailureResponse() models.ResponseCode {
	return models.FailureResponseCodes[0]
}

func constructInfoResponse() models.ResponseCode {
	return models.InformationResponseCode[0]
}

func constructFraudResponse() models.ResponseCode {
	return models.FraudResponseCode[0]
}
