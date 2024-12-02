package models

import "liquid/pkg/utils"

func (assert *ModelsTestSuite) TestGetCurrencies() {
	// expect currencies to be populated
	assert.LessOrEqual(3, len(GetCurrencies()))
	assert.Contains(utils.Map(GetCurrencies(), func(c Currency) CurrencyName { return c.Name }), CurrencyName("USD"))
}
