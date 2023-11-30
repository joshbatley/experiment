package event

import (
	utils "shared"
	"time"
)

func New(clientId string, paymentId string) *Event {
	id := utils.NewEventId()
	if paymentId == "" {
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
	ne.Action = ActionRequested
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

func (e *Event) AsCapture(maxCapturableAmount int) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = ResponseCodeSuccess
	ne.Action = ActionCapture

	if utils.RandomChance(10) && maxCapturableAmount > 0 {
		ne.CapturedAmount = utils.GenerateRandomNumberBetween(maxCapturableAmount)
	} else {
		ne.CapturedAmount = maxCapturableAmount
	}

	if ne.CapturedAmount == maxCapturableAmount {
		ne.Status = StatusCaptured
	} else {
		ne.Status = StatusPartiallyCaptured
	}
	return ne
}

func (e *Event) AsRefund(maxRefundableAmount int, isFullyCaptured bool) *Event {
	ne := New(e.ClientId, e.PaymentID)
	ne.ResponseCode = ResponseCodeSuccess
	ne.Action = ActionRefund

	if utils.RandomChance(10) && maxRefundableAmount > 0 {
		ne.RefundedAmount = utils.GenerateRandomNumberBetween(maxRefundableAmount)
	} else {
		ne.RefundedAmount = maxRefundableAmount
	}
	if ne.RefundedAmount == maxRefundableAmount && isFullyCaptured {
		ne.Status = StatusRefunded
	} else {
		ne.Status = StatusPartiallyRefunded
	}
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
