package models

func (assert *ModelsTestSuite) TestCreatePaymentAccount() {
	// given an application entity
	res := CreateClientEntity("test")

	// should return tigerbeetle id
	id := CreatePaymentAccount(res, CurrencyName("USD"))
	assert.NotNil(id)

	paymentAccount := FindPaymentAccountByAppEntityIdAndCurrencyName(res, CurrencyName("USD"))
	assert.NotNil(paymentAccount)
}

func (assert *ModelsTestSuite) TestEmptyPaymentAccount() {
	// given an application entity
	res := CreateClientEntity("test")

	// should be nil
	paymentAccount := FindPaymentAccountByAppEntityIdAndCurrencyName(res, CurrencyName("EUR"))

	assert.Nil(paymentAccount)
}
