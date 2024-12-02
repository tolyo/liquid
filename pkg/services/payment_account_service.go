package services

import (
	"liquid/pkg/models"
	"liquid/pkg/tiger"
	"liquid/pkg/utils"
	"math/big"

	log "github.com/sirupsen/logrus"
	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func GetFullPaymentAccounts(externalId models.AppEntityExternalId) []models.FullPaymentAccount {

	paymentAccounts := models.FindAllPaymentAccountsByAppEntityExternalId(externalId)
	accountIds := utils.Map(paymentAccounts, func(pa models.PaymentAccount) tigerTypes.Uint128 {
		return tigerTypes.ToUint128(uint64(pa.TigerId))
	})

	// Check the sums for both accounts
	tbAccounts, err := tiger.Instance().LookupAccounts(accountIds)
	if err != nil {
		log.Fatalf("Could not fetch accounts: %s", err)
	}

	log.Info(tbAccounts)

	fullPaymentAccount := make([]models.FullPaymentAccount, len(tbAccounts))

	for index, tbAccount := range tbAccounts {
		amount := new(big.Int)
		debits := tbAccount.DebitsPosted.BigInt()
		credits := tbAccount.CreditsPosted.BigInt()
		amount.Sub(&debits, &credits)
		fullPaymentAccount[index] = models.FullPaymentAccount{
			PaymentAccount: paymentAccounts[index],
			Amount:         amount.String(),
		}
	}

	return fullPaymentAccount
}
