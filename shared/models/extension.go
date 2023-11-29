package models

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"reflect"
	utils "shared"
	"strings"
	"time"
)

func NewEvent(clientId string) *Event {
	id := utils.NewEventId()
	return &Event{
		ID:        id,
		PaymentID: utils.NewPaymentId(),
		ActionID:  utils.NewActionId(id),
		ClientId:  clientId,
		Reference: utils.GenerateRandomReference(8),
		Timestamp: time.Now(),
	}
}

func randomChance() bool {
	randomNum := rand.Intn(10)
	return randomNum == 0
}

func (e *Event) updateIds() {
	id := utils.NewEventId()
	e.ID = id
	e.ActionID = utils.NewActionId(id)
}

func (e *Event) clearEventSpecific() {

}

func (e *Event) AsRequested() *Event {
	e.updateIds()
	e.Action = ActionRequested
	e.AuthorizedAmount = utils.GenerateRandomNumber()
	e.Status = StatusPending
	return e
}

func (e *Event) AsAuthorized() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionAuthorize
	e.Status = StatusAuthorized
	return e
}

func (e *Event) AsCapture() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionCapture

	if randomChance() {
		e.CapturedAmount = utils.GenerateRandomNumberBetween(e.AuthorizedAmount - e.CapturedAmount)
	} else {
		e.CapturedAmount = e.AuthorizedAmount
	}
	if e.CapturedAmount == e.AuthorizedAmount {
		e.Status = StatusCaptured
	} else {
		e.Status = StatusPartiallyCaptured
	}
	return e
}

func (e *Event) AsRefund() *Event {
	e.updateIds()
	e.ResponseCode = ResponseCodeSuccess
	e.Action = ActionRefund

	if randomChance() {
		e.RefundedAmount = utils.GenerateRandomNumberBetween(e.CapturedAmount - e.RefundedAmount)
	} else {
		e.RefundedAmount = e.CapturedAmount
	}
	if e.RefundedAmount == e.CapturedAmount {
		e.Status = StatusRefunded
	} else {
		e.Status = StatusPartiallyRefunded
	}
	return e
}

func (e *Event) AsVoid() *Event {
	e.updateIds()
	e.Action = ActionVoid
	e.ResponseCode = ResponseCodeSuccess
	e.Status = StatusCancelled
	return e
}

func (e *Event) AsExpiry() *Event {
	e.updateIds()
	e.Action = ActionExpiry
	e.ResponseCode = ResponseCodeSuccess
	return e
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

// Move out
const (
	LoggerTag = "log"
	Omit      = "omitempty"
)

func LogTags(obj interface{}) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	l := log.Info()
	q := objType.NumField()
	for i := 0; i < q; i++ {
		f := objType.Field(i)
		tag := f.Tag.Get(LoggerTag)
		v := objValue.Field(i)
		n := f.Name
		switch f.Type.Kind() {
		case reflect.String:
			val := v.String()
			if strings.Contains(tag, Omit) && val == "" {
				continue
			}
			l.Str(n, val)
		case reflect.Float64:
			val := v.Float()
			if strings.Contains(tag, Omit) && val == 0 {
				continue
			}
			l.Float64(n, val)
		default:
			val := v.Interface()
			if strings.Contains(tag, Omit) && isStructEmpty(val) {
				continue
			}
			l.Interface(n, val)
		}

	}
	l.Send()
}

func isStructEmpty(s interface{}) bool {
	sValue := reflect.ValueOf(s)
	emptyValue := reflect.New(sValue.Type()).Elem()

	return reflect.DeepEqual(sValue.Interface(), emptyValue.Interface())
}
