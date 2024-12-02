package services

import (
	"liquid/pkg/models"
	"liquid/pkg/tiger"
	"liquid/pkg/utils"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func CreateExchange(
	externalId models.AppEntityExternalId,
	amount int,
	currencySell models.CurrencyName,
	currencyBuy models.CurrencyName,
) (tigerTypes.CreateTransferResult, error) {
	log.Infof("CreateExchange initiated for externalId: %s, Amount: %d, Currency Sell: %s, Currency Buy: %s", externalId, amount, currencySell, currencyBuy)

	currencyInstanceSell := models.GetCurrency(currencySell)
	currencyInstanceBuy := models.GetCurrency(currencyBuy)
	log.Infof("Currency instance for Sell: %+v, Currency instance for Buy: %+v", currencyInstanceSell, currencyInstanceBuy)

	masterAccounts := models.FindAllPaymentAccountsByAppEntityExternalId("MASTER")
	clientAccounts := models.FindAllPaymentAccountsByAppEntityExternalId(externalId)

	log.Infof("Found %d master accounts and %d client accounts for externalId: %s", len(masterAccounts), len(clientAccounts), externalId)

	clientDebitAccountTigerId := utils.Filter(clientAccounts, func(x models.PaymentAccount) bool {
		return x.Currency == currencySell
	})[0].TigerId

	masterCreditAccountTigerId := utils.Filter(masterAccounts, func(x models.PaymentAccount) bool {
		return x.Currency == currencySell
	})[0].TigerId

	masterDebitAccountTigerId := utils.Filter(masterAccounts, func(x models.PaymentAccount) bool {
		return x.Currency == currencyBuy
	})[0].TigerId

	clientCreditAccountTigerId := utils.Filter(clientAccounts, func(x models.PaymentAccount) bool {
		return x.Currency == currencyBuy
	})[0].TigerId

	log.Infof("Client Debit Account TigerId: %d, Master Credit Account TigerId: %d, Master Debit Account TigerId: %d, Client Credit Account TigerId: %d",
		clientDebitAccountTigerId, masterCreditAccountTigerId, masterDebitAccountTigerId, clientCreditAccountTigerId)

	currencyPair := models.CurrencyPair(currencyBuy + "/" + currencySell)
	amountBuy := decimal.NewFromUint64(uint64(amount)).Div(FxRates[currencyPair]).RoundBank(0)
	fee := models.GetFee(currencyPair)
	feeAmount := decimal.NewFromUint64(uint64(amount)).Mul(fee.Percent).RoundBank(0)

	log.Infof("Currency Pair: %s, FX Rate: %s, Amount to sell: %d, Fee Percent: %d, Fee Amount: %d", currencyPair, FxRates[currencyPair].String(), amount, fee.Percent, feeAmount)

	transfers := []tigerTypes.Transfer{
		{
			ID:              tigerTypes.ID(),
			DebitAccountID:  tigerTypes.ToUint128(uint64(masterCreditAccountTigerId)),
			CreditAccountID: tigerTypes.ToUint128(uint64(clientDebitAccountTigerId)),
			Amount:          tigerTypes.ToUint128(uint64(amount)),
			Ledger:          uint32(currencyInstanceSell.Ledger),
			Code:            1,
			Flags: tigerTypes.TransferFlags{
				Linked: true,
			}.ToUint16(),
			Timestamp: 0,
		},
		{
			ID:              tigerTypes.ID(),
			DebitAccountID:  tigerTypes.ToUint128(uint64(clientCreditAccountTigerId)),
			CreditAccountID: tigerTypes.ToUint128(uint64(masterDebitAccountTigerId)),
			Amount:          tigerTypes.ToUint128(amountBuy.BigInt().Uint64()),
			Ledger:          uint32(currencyInstanceBuy.Ledger),
			Code:            1,
			Flags: tigerTypes.TransferFlags{
				Linked: true,
			}.ToUint16(),
			Timestamp: 0,
		},
		{
			ID:              tigerTypes.ID(),
			DebitAccountID:  tigerTypes.ToUint128(uint64(masterCreditAccountTigerId)),
			CreditAccountID: tigerTypes.ToUint128(uint64(clientDebitAccountTigerId)),
			Amount:          tigerTypes.ToUint128(feeAmount.BigInt().Uint64()),
			Ledger:          uint32(currencyInstanceSell.Ledger),
			Code:            1,
			Flags:           0,
			Timestamp:       0,
		},
	}

	for _, value := range transfers {
		log.Infof("Details: Amount: %s, DebitAccountID: %s, CreditAccountID: %s, Ledger: %d, Flags: %d, Timestamp: %d",
			value.Amount.String(), value.DebitAccountID.String(), value.CreditAccountID.String(), value.Ledger, value.Flags, value.Timestamp)
	}

	transferRes, err := tiger.Instance().CreateTransfers(transfers)
	if err != nil {
		log.Errorf("Failed to create transfers for externalId: %s, Error: %v", externalId, err)
		return tigerTypes.TransferOK, err
	}

	var successful = true
	var result = tigerTypes.TransferOK

	for _, res := range transferRes {
		if res.Result != tigerTypes.TransferOK {
			successful = false
			result = res.Result
			log.Errorf("Transfer failed : %+v", transferRes)
			break
		}
	}
	if successful {
		SendPaymentEvent(currencyBuy)
		log.Infof("Exchange creation completed for externalId: %s, Amount: %d, Currency Buy: %s, Currency Sell: %s", externalId, amount, currencyBuy, currencySell)
	}
	return result, nil

}
