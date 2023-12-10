package event

type PaymentMethod string

const (
	PaymentMethodPayPal          PaymentMethod = "PayPal"
	PaymentMethodBankTransfer    PaymentMethod = "Bank Transfer"
	PaymentMethodApplePay        PaymentMethod = "Apple Pay"
	PaymentMethodGooglePay       PaymentMethod = "Google Pay"
	PaymentMethodStripe          PaymentMethod = "Stripe"
	PaymentMethodVenmo           PaymentMethod = "Venmo"
	PaymentMethodVisa            PaymentMethod = "Visa"
	PaymentMethodMastercard      PaymentMethod = "Mastercard"
	PaymentMethodAmericanExpress PaymentMethod = "American Express"
	PaymentMethodDiscover        PaymentMethod = "Discover"
	PaymentMethodJCB             PaymentMethod = "JCB"
	PaymentMethodDinersClub      PaymentMethod = "Diners Club"
)

var paymentMethods = []PaymentMethod{
	PaymentMethodPayPal, PaymentMethodBankTransfer, PaymentMethodApplePay, PaymentMethodGooglePay, PaymentMethodStripe,
	PaymentMethodVenmo, PaymentMethodVisa, PaymentMethodMastercard, PaymentMethodAmericanExpress, PaymentMethodDiscover,
	PaymentMethodJCB, PaymentMethodDinersClub,
}
