package models

import "github.com/shopspring/decimal"

func (assert *ModelsTestSuite) TestGetFees() {
	// expect currencies to be populated
	assert.NotNil(GetFee(CurrencyPair("EUR/USD")))
	expected, _ := decimal.NewFromString("0.02")
	assert.Equal(expected, GetFee(CurrencyPair("EUR/USD")).Percent)
	assert.Equal(CurrencyPair("EUR/USD"), GetFee(CurrencyPair("EUR/USD")).Currency_pair)
}
