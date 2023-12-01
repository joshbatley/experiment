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

func (e *Event) AsRequested() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionRequest
	ne.AuthorizedAmount = utils.GenerateRandomNumber()
	ne.Status = StatusPending
	return ne
}

func (e *Event) AsAuthorized() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = ResponseCodeSuccess
	ne.Action = ActionAuthorize
	ne.Status = StatusAuthorized
	ne.AuthorizedAmount = e.AuthorizedAmount
	return ne
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

func (e *Event) AsCapture(maxCapturableAmount int) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = ResponseCodeSuccess
	ne.Action = ActionCapture
	ne.CapturedAmount = setCapturedAmount(maxCapturableAmount)
	ne.Status = setCaptureStatus(ne.CapturedAmount, maxCapturableAmount)
	return ne
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

func (e *Event) AsRefund(maxRefundableAmount int, isFullyCaptured bool) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = ResponseCodeSuccess
	ne.Action = ActionRefund

	ne.RefundedAmount = setRefundedAmount(maxRefundableAmount)
	ne.Status = setRefundStatus(ne.RefundedAmount, maxRefundableAmount, isFullyCaptured)
	return ne
}

func (e *Event) AsVoid() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionVoid
	ne.ResponseCode = ResponseCodeSuccess
	ne.Status = StatusCancelled
	return ne
}

func (e *Event) AsExpiry() *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.Action = ActionExpiry
	ne.ResponseCode = ResponseCodeSuccess
	return ne
}

func (e *Event) withCustomer(customer Customer) *Event {
	e.Customer = customer
	return e
}

func (e *Event) withRecipient(recipient Recipient) *Event {
	e.Recipient = recipient
	return e
}

func (e *Event) withShipping(address Address) *Event {
	e.ShippingAddress = address
	return e
}

func (e *Event) withBilling(address Address) *Event {
	e.BillingAddress = address
	return e
}

func (e *Event) withCardDetails(details CardDetails) *Event {
	e.CardDetails = details
	return e
}

func (e *Event) withItems(items ...Item) *Event {
	e.Items = items
	return e
}

func (e *Event) withPayment(currency CurrencyCode, paymentMethod PaymentMethod) *Event {
	e.Currency = currency
	e.PaymentMethod = paymentMethod
	return e
}
