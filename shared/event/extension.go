package event

import (
	utils "shared"
	"time"
)

func New(clientId string, paymentId string) *Event {
	id := utils.NewEventId()
	if len(paymentId) == 0 {
		paymentId = utils.NewPaymentId()
	}
	return &Event{
		ID:        id,
		PaymentID: paymentId,
		ActionID:  utils.NewActionId(id),
		ClientId:  clientId,
		Reference: utils.GenerateRandomReference(8),
		Timestamp: time.Now(),
	}
}

func getRandomSuccessCode() ResponseCode {
	return utils.GetRandomItem(successfulResponseCodes)
}

func getRandomFailureCode() ResponseCode {
	return utils.GetRandomItem(failureResponseCodes)
}

func setCaptureStatus(capture int, maxCapturableAmount int) Status {
	if capture == maxCapturableAmount {
		return StatusCaptured
	}
	return StatusPartiallyCaptured
}

func setCapturedAmount(maxCapturableAmount int) int {
	if utils.RandomChance(10) && maxCapturableAmount > 0 {
		return utils.GenerateRandomNumberBetween(maxCapturableAmount)
	}
	return maxCapturableAmount
}

func setRefundStatus(refund int, maxRefundableAmount int, isFullyCaptured bool) Status {
	if refund == maxRefundableAmount && isFullyCaptured {
		return StatusRefunded
	}
	return StatusPartiallyRefunded
}

func setRefundedAmount(maxRefundableAmount int) int {
	if utils.RandomChance(10) && maxRefundableAmount > 0 {
		return utils.GenerateRandomNumberBetween(maxRefundableAmount)
	}
	return maxRefundableAmount
}

func (e *Event) withRandomChanceOfFailure() *Event {
	if !utils.RandomChance(100) {
		return e
	}
	e.ResponseCode = getRandomFailureCode()
	e.Status = StatusFailed
	return e
}

func (e *Event) withCustomer() *Event {
	e.Customer = Customer{
		PhoneNumber: utils.GetRandomItem(fakePhone),
		Email:       utils.GetRandomItem(fakeEmail),
		Name:        utils.GetRandomItem(fakeName),
	}
	return e
}

func (e *Event) withRecipient() *Event {
	e.Recipient = Recipient{
		Name:          utils.GetRandomItem(fakeName),
		AccountNumber: utils.GetRandomItem(fakeAccountNo),
		BankName:      utils.GetRandomItem(fakeBank),
	}
	return e
}

func (e *Event) withShipping() *Event {
	e.ShippingAddress = Address{
		Street1:    utils.GetRandomItem(fakeStreet1),
		Street2:    utils.GetRandomItem(fakeStreet2),
		City:       utils.GetRandomItem(fakeCity),
		State:      utils.GetRandomItem(fakeState),
		PostalCode: utils.GetRandomItem(fakePostalCode),
		Country:    utils.GetRandomItem(fakeCountry),
	}
	return e
}

func (e *Event) withBilling() *Event {
	e.BillingAddress = Address{
		Street1:    utils.GetRandomItem(fakeStreet1),
		Street2:    utils.GetRandomItem(fakeStreet2),
		City:       utils.GetRandomItem(fakeCity),
		State:      utils.GetRandomItem(fakeState),
		PostalCode: utils.GetRandomItem(fakePostalCode),
		Country:    utils.GetRandomItem(fakeCountry),
	}
	return e
}

func (e *Event) withCardDetails() *Event {
	month := utils.GenerateRandomNumberBetween(12)
	if month == 0 {
		month++
	}
	year := utils.GenerateRandomNumberBetween(99)
	if year < 23 {
		year = 23
	}

	e.CardDetails = CardDetails{
		CardNumber:     utils.GetRandomItem(fakeCardNumber),
		ExpiryMonth:    month,
		ExpiryYear:     year,
		CVV:            utils.GetRandomItem(fakeCvv),
		CardholderName: utils.GetRandomItem(fakeName),
		IssuingBank:    utils.GetRandomItem(fakeBank),
	}
	return e
}

func (e *Event) withCurrency() *Event {
	e.Currency = utils.GetRandomItem(currencies)
	return e
}

func (e *Event) withPaymentMethod() *Event {
	e.PaymentMethod = utils.GetRandomItem(paymentMethods)
	return e
}

func (e *Event) AsRequested() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionRequest
	ne.AuthorizedAmount = utils.GenerateRandomNumber()
	ne.Status = StatusPending
	ne.ResponseCode = getRandomSuccessCode()
	return ne.withPaymentMethod().withCurrency().withCardDetails().withBilling().withCustomer().withRecipient().withShipping().withRandomChanceOfFailure()
}

func (e *Event) AsAuthorized() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = getRandomSuccessCode()
	ne.Action = ActionAuthorize
	ne.Status = StatusAuthorized
	ne.AuthorizedAmount = e.AuthorizedAmount
	return ne.withRandomChanceOfFailure()
}

func (e *Event) AsCapture(maxCapturableAmount int) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = getRandomSuccessCode()
	ne.Action = ActionCapture
	ne.CapturedAmount = setCapturedAmount(maxCapturableAmount)
	ne.Status = setCaptureStatus(ne.CapturedAmount, maxCapturableAmount)
	return ne.withRandomChanceOfFailure()
}

func (e *Event) AsRefund(maxRefundableAmount int, isFullyCaptured bool) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = getRandomSuccessCode()
	ne.Action = ActionRefund

	ne.RefundedAmount = setRefundedAmount(maxRefundableAmount)
	ne.Status = setRefundStatus(ne.RefundedAmount, maxRefundableAmount, isFullyCaptured)
	return ne.withRandomChanceOfFailure()
}

func (e *Event) AsVoid() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionVoid
	ne.ResponseCode = getRandomFailureCode()
	ne.Status = StatusCancelled
	return ne
}

func (e *Event) AsExpiry() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionExpiry
	ne.ResponseCode = getRandomFailureCode()
	ne.Status = StatusFailed
	return ne
}

func (e *Event) IsSuccessfulResponseCode() bool {
	return utils.FindIndex(successfulResponseCodes, e.ResponseCode) > 0
}

func (e *Event) IsFailureResponseCode() bool {
	return utils.FindIndex(failureResponseCodes, e.ResponseCode) > 0
}
