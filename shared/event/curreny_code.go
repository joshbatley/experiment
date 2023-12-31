package event

type CurrencyCode string

const (
	CurrencyUSD CurrencyCode = "USD"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyJPY CurrencyCode = "JPY"
	CurrencyGBP CurrencyCode = "GBP"
	CurrencyAUD CurrencyCode = "AUD"
	CurrencyCAD CurrencyCode = "CAD"
	CurrencyCHF CurrencyCode = "CHF"
	CurrencyCNY CurrencyCode = "CNY"
	CurrencySEK CurrencyCode = "SEK"
	CurrencyNZD CurrencyCode = "NZD"
	CurrencyKRW CurrencyCode = "KRW"
	CurrencyHKD CurrencyCode = "HKD"
	CurrencyNOK CurrencyCode = "NOK"
	CurrencyMXN CurrencyCode = "MXN"
	CurrencyDKK CurrencyCode = "DKK"
	CurrencySGD CurrencyCode = "SGD"
	CurrencyPLN CurrencyCode = "PLN"
	CurrencyINR CurrencyCode = "INR"
	CurrencyRUB CurrencyCode = "RUB"
	CurrencyBRL CurrencyCode = "BRL"
)

var currencies = []CurrencyCode{
	CurrencyUSD, CurrencyEUR, CurrencyJPY, CurrencyGBP, CurrencyAUD, CurrencyCAD, CurrencyCHF, CurrencyCNY,
	CurrencySEK, CurrencyNZD, CurrencyKRW, CurrencyHKD, CurrencyNOK, CurrencyMXN, CurrencyDKK, CurrencySGD,
	CurrencyPLN, CurrencyINR, CurrencyRUB, CurrencyBRL,
}
