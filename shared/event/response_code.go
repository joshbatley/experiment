package event

import utils "shared"

type ResponseCode string

const (
	// Success category
	ResponseCodeSuccessfulAction ResponseCode = "2XX"
	ResponseCodeSuccess          ResponseCode = "200"

	// Failure category
	ResponseCodeFailure              ResponseCode = "4XX"
	ResponseCodeInvalidData          ResponseCode = "400"
	ResponseCodeInsufficientFunds    ResponseCode = "401"
	ResponseCodeDuplicateTransaction ResponseCode = "402"
	ResponseCodeCustomFailureCode1   ResponseCode = "403"
	ResponseCodeCustomFailureCode2   ResponseCode = "404"

	// Information category
	ResponseCodeExtraInformation ResponseCode = "3XX"
	ResponseCodeAdditionalData   ResponseCode = "301"
	ResponseCodePendingAction    ResponseCode = "302"
	ResponseCodeCustomCode1      ResponseCode = "303"
	ResponseCodeCustomCode2      ResponseCode = "304"

	// Fraud category
	ResponseCodeFraudDetected      ResponseCode = "6XX"
	ResponseCodeSuspiciousActivity ResponseCode = "601"
	ResponseCodeIdentityTheft      ResponseCode = "602"
	ResponseCodeCustomFraudCode1   ResponseCode = "603"
	ResponseCodeCustomFraudCode2   ResponseCode = "604"
)

var successfulResponseCodes = []ResponseCode{
	ResponseCodeSuccessfulAction,
	ResponseCodeSuccess,
}

var failureResponseCodes = []ResponseCode{
	ResponseCodeFailure,
	ResponseCodeInvalidData,
	ResponseCodeInsufficientFunds,
	ResponseCodeDuplicateTransaction,
	ResponseCodeCustomFailureCode1,
	ResponseCodeCustomFailureCode2,
}

var informationResponseCode = []ResponseCode{
	ResponseCodeExtraInformation,
	ResponseCodeAdditionalData,
	ResponseCodePendingAction,
	ResponseCodeCustomCode1,
	ResponseCodeCustomCode2,
}

var fraudResponseCode = []ResponseCode{
	ResponseCodeFraudDetected,
	ResponseCodeSuspiciousActivity,
	ResponseCodeIdentityTheft,
	ResponseCodeCustomFraudCode1,
	ResponseCodeCustomFraudCode2,
}

func GetRandomSuccessCode() ResponseCode {
	return utils.GetRandomItem(successfulResponseCodes)
}

func GetRandomFailureCode() ResponseCode {
	return utils.GetRandomItem(failureResponseCodes)
}

func GetRandomInfoCode() ResponseCode {
	return utils.GetRandomItem(informationResponseCode)
}

func GetRandomFraudCode() ResponseCode {
	return utils.GetRandomItem(fraudResponseCode)
}
