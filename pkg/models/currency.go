package models

import (
	"liquid/pkg/db"

	"github.com/shopspring/decimal"
)

type CurrencyPair string

type CurrencyName string

type CurrencyPrecision int

type CurrencyLedger int

type Currency struct {
	Name       CurrencyName
	Precision  CurrencyPrecision
	Ledger     CurrencyLedger
	Min_volume decimal.Decimal
	Max_volume decimal.Decimal
}

const (
	USD CurrencyName = "USD"
	EUR CurrencyName = "EUR"
	JPY CurrencyName = "JPY"
	GBP CurrencyName = "GBP"
	AUD CurrencyName = "AUD"
)

func GetCurrencies() []Currency {
	res := db.QueryList[Currency](`SELECT * FROM currency`)
	return res
}

func GetCurrency(name CurrencyName) Currency {
	res := db.QueryVal[Currency](`SELECT * FROM currency WHERE name = $1`, string(name))
	return res
}
