package models

import (
	"liquid/pkg/db"

	"github.com/shopspring/decimal"
)

type Fee struct {
	Currency_pair CurrencyPair
	Percent       decimal.Decimal
}

func GetFee(currencyPair CurrencyPair) Fee {
	res := db.QueryVal[Fee](`SELECT * FROM fee WHERE currency_pair = $1`, string(currencyPair))
	return res
}
