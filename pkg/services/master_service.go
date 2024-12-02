package services

import (
	"liquid/pkg/models"
	"liquid/pkg/tiger"
	"liquid/pkg/utils"
	"math/big"

	log "github.com/sirupsen/logrus"

	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

var defaultAmounts = map[string]int{
	string(models.USD): 1000000,
	string(models.EUR): 921658,
	string(models.JPY): 109890110,
	string(models.GBP): 750000,
	string(models.AUD): 1349528,
}

func InitProviderAccounts() {
	log.Info("Initializing provider accounts")
	providerId := models.GetProvider()

	for _, c := range models.GetCurrencies() {
		paymentAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(providerId, c.Name)

		if paymentAccount == nil {
			// create the respective account in DB
			tigerId := models.CreatePaymentAccount(providerId, c.Name)
			log.Infof("Created payment account with tigerId: %d", tigerId)
			res, err := tiger.Instance().CreateAccounts([]tigerTypes.Account{
				{
					ID:     tigerTypes.ToUint128(uint64(tigerId)),
					Ledger: uint32(c.Ledger),
					Code:   1000,
					Flags: tigerTypes.AccountFlags{
						DebitsMustNotExceedCredits: true,
					}.ToUint16(),
				},
			})
			if err != nil {
				log.Fatalf("Error creating account in TigerBeetle for %s: %v", c.Name, err)
			}

			log.Infof("Account creation result for %s: %+v", c.Name, res)
		}
	}
}

func InitMasterAccounts() {
	log.Info("Initializing master accounts")
	masterId := models.GetMaster()

	for _, c := range models.GetCurrencies() {
		paymentAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(masterId, c.Name)
		if paymentAccount == nil {
			tigerId := models.CreatePaymentAccount(masterId, c.Name)
			log.Infof("Created payment account with tigerId: %d", tigerId)
			res, err := tiger.Instance().CreateAccounts([]tigerTypes.Account{
				{
					ID:     tigerTypes.ToUint128(uint64(tigerId)),
					Ledger: uint32(c.Ledger),
					Code:   1001,
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

			// get provider
			providerTigerId := models.FindPaymentAccountByAppEntityIdAndCurrencyName(models.GetProvider(), c.Name).TigerId

			// credit with initial amounts
			transfers := []tigerTypes.Transfer{{
				ID:              tigerTypes.ID(), // TigerBeetle time-based ID.
				DebitAccountID:  tigerTypes.ToUint128(uint64(tigerId)),
				CreditAccountID: tigerTypes.ToUint128(uint64(providerTigerId)),
				Amount:          tigerTypes.ToUint128(uint64(defaultAmounts[string(c.Name)] * 100)),
				Ledger:          uint32(c.Ledger),
				Code:            1,
				Flags:           0,
				Timestamp:       0,
			}}

			transferRes, err := tiger.Instance().CreateTransfers(transfers)
			if err != nil {
				log.Fatalf("Error transferring initial amount for %s: %v", c.Name, err)
			}

			log.Infof("Initial transfer for %s completed: %+v", c.Name, transferRes)

		}
	}
}

func InitTestAccount(externalId string) {
	log.Info("Init test account:  " + externalId)
	// Get seed
	_, err := models.FindAppEntityExternalId(models.AppEntityExternalId(externalId))
	if err == nil {
		log.Info("User already exists")
		return
	}
	log.Infof("Creating client entity externa_id: %s", externalId)
	CreateClientEntityWithAccounts(models.AppEntityExternalId(externalId))

	for _, c := range models.GetCurrencies() {
		tigerId := models.FindPaymentAccountByAppEntityExternalIdAndCurrencyName(models.AppEntityExternalId(externalId), c.Name).TigerId
		providerTigerId := models.FindPaymentAccountByAppEntityIdAndCurrencyName(models.GetProvider(), c.Name).TigerId

		// credit with initial amounts
		transfers := []tigerTypes.Transfer{{
			ID:              tigerTypes.ID(), // TigerBeetle time-based ID.
			DebitAccountID:  tigerTypes.ToUint128(uint64(tigerId)),
			CreditAccountID: tigerTypes.ToUint128(uint64(providerTigerId)),
			Amount:          tigerTypes.ToUint128(uint64(100000)),
			Ledger:          uint32(c.Ledger),
			Code:            1,
			Flags:           0,
			Timestamp:       0,
		}}

		transferRes, err := tiger.Instance().CreateTransfers(transfers)
		if err != nil {
			log.Fatal("Unable to create seed transfers for test account")
		}
		log.Infof("Initial transfer for %s completed: %+v", c.Name, transferRes)
	}
}

// GetMasterAccounts returns all master accounts with their balances.
func GetMasterAccounts() []models.FullPaymentAccount {
	masterId := models.GetMaster()
	paymentAccounts := models.FindAllPaymentAccountsByAppEntityId(masterId)

	accountIds := utils.Map(paymentAccounts, func(pa models.PaymentAccount) tigerTypes.Uint128 {
		return tigerTypes.ToUint128(uint64(pa.TigerId))
	})
	// Check the sums for both accounts
	tbAccounts, err := tiger.Instance().LookupAccounts(accountIds)
	if err != nil {
		log.Fatalf("Error finding master accounts in tigerbeetle: %s", err)
	}

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

// GetMasterAccountForCurrency returns the master account for a specific currency.
func GetMasterAccountForCurrency(currency models.CurrencyName) models.FullPaymentAccount {
	return utils.Filter(GetMasterAccounts(), func(account models.FullPaymentAccount) bool {
		return account.Currency == currency
	})[0]
}
