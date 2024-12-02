package services

import (
	"liquid/pkg/models"
	"liquid/pkg/tiger"

	log "github.com/sirupsen/logrus"
	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func CreateClientEntityWithAccounts(id models.AppEntityExternalId) {
	clientEntity := models.CreateClientEntity(id)
	for _, c := range models.GetCurrencies() {
		id := models.CreatePaymentAccount(clientEntity, c.Name)
		res, err := tiger.Instance().CreateAccounts([]tigerTypes.Account{
			{
				ID:     tigerTypes.ToUint128(uint64(id)),
				Ledger: uint32(c.Ledger),
				Code:   1002,
				Flags: tigerTypes.AccountFlags{
					History:                    true,
					CreditsMustNotExceedDebits: true,
				}.ToUint16(),
			},
		})
		if err != nil {
			log.Fatalf("Error creating account in TigerBeetle for %s: %v", c.Name, err)
		}

		log.Infof("Account creation result for %s: %+v", c.Name, res)
	}
}
