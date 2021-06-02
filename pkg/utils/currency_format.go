package utils

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const Divisor = 100.0

func CurrencyFormat(currencyCode string, amount int64) string {
	tag := getCurrencyTag(currencyCode)
	cur, _ := currency.FromTag(tag)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(float64(amount)/Divisor, number.Scale(scale))
	p := message.NewPrinter(tag)
	return p.Sprintf("%v %#v", currency.Symbol(cur), dec)
}

func getCurrencyTag(s string) language.Tag {
	switch s {
	case "USD":
		return language.MustParse("en")
	case "EUR":
		return language.MustParse("eu")
	default:
		return language.MustParse("uk")
	}
}
