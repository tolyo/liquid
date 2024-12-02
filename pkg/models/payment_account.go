package models

import (
	"liquid/pkg/db"
	"log"
)

type PaymentAccountTigerId uint

type PaymentAccount struct {
	AppEntityId AppEntityId           `json:"-"`
	TigerId     PaymentAccountTigerId `json:"-"`
	Currency    CurrencyName          `json:"currency"`
}

type FullPaymentAccount struct {
	PaymentAccount
	Amount string `json:"amount"`
}

const basePaymentAccountQuery = `
	SELECT
	  ae.pub_id,
	  pa.tigerbeetle_id,
	  c.name

	FROM payment_account AS pa

	INNER JOIN app_entity ae
	  ON pa.app_entity_id = ae.id

	INNER JOIN currency c
	  ON pa.currency_name = c.name
  `

func CreatePaymentAccount(clientEntityId AppEntityId, currencyName CurrencyName) PaymentAccountTigerId {
	var id int
	err := db.Instance().QueryRow("SELECT create_payment_account($1, $2)", clientEntityId, currencyName).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return PaymentAccountTigerId(id)
}

func FindPaymentAccountByAppEntityIdAndCurrencyName(
	appEntityId AppEntityId,
	currencyName CurrencyName,
) *PaymentAccount {
	var paymentAccount PaymentAccount
	err := db.Instance().QueryRow(
		basePaymentAccountQuery+`WHERE ae.pub_id = $1 AND c.name = $2`,
		appEntityId,
		currencyName,
	).Scan(
		&paymentAccount.AppEntityId,
		&paymentAccount.TigerId,
		&paymentAccount.Currency,
	)
	if err != nil {
		return nil
	}
	return &paymentAccount
}

func FindAllPaymentAccountsByAppEntityId(
	appEntityId AppEntityId,
) []PaymentAccount {

	res := make([]PaymentAccount, 0)

	rows, err := db.Instance().Query(basePaymentAccountQuery+`WHERE ae.pub_id = $1`, appEntityId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var paymentAccount PaymentAccount
		err := rows.Scan(&paymentAccount.AppEntityId,
			&paymentAccount.TigerId,
			&paymentAccount.Currency)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, paymentAccount)
	}
	return res
}

func FindPaymentAccountByAppEntityExternalIdAndCurrencyName(
	appEntityExternalId AppEntityExternalId,
	currencyName CurrencyName,
) *PaymentAccount {
	var paymentAccount PaymentAccount
	err := db.Instance().QueryRow(
		basePaymentAccountQuery+`WHERE ae.external_id = $1 AND c.name = $2`,
		appEntityExternalId,
		currencyName,
	).Scan(
		&paymentAccount.AppEntityId,
		&paymentAccount.TigerId,
		&paymentAccount.Currency,
	)
	if err != nil {
		return nil
	}
	return &paymentAccount
}

func FindAllPaymentAccountsByAppEntityExternalId(
	appEntityExternalId AppEntityExternalId,
) []PaymentAccount {

	res := make([]PaymentAccount, 0)

	rows, err := db.Instance().Query(basePaymentAccountQuery+`WHERE ae.external_id = $1`, appEntityExternalId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var paymentAccount PaymentAccount
		err := rows.Scan(&paymentAccount.AppEntityId,
			&paymentAccount.TigerId,
			&paymentAccount.Currency)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, paymentAccount)
	}
	return res
}
