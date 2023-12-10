package event

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
