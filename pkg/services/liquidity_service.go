package services

import (
	"liquid/pkg/conf"
	"liquid/pkg/models"
	"liquid/pkg/tiger"
	"time"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

var mockDelay = map[string]int{
	"USD": 3,
	"EUR": 2,
	"JPY": 3,
	"GBP": 2,
	"AUD": 3,
}

var paymentChannel chan PaymentEvent

type PaymentEvent struct {
	Currency models.CurrencyName
}

func InitPaymentChannel() {
	paymentChannel = make(chan PaymentEvent, 100)
	log.Info("Payment channel initialized")
}

func ListenForPayments() {
	for event := range paymentChannel {
		log.Infof("Received payment event for currency: %s", event.Currency)
		currencyInstance := models.GetCurrency(event.Currency)

		masterAccount := GetMasterAccountForCurrency(event.Currency)
		masterAmount, _ := decimal.NewFromString(masterAccount.Amount)

		log.Infof("Checking liquidity for currency: %s, master amount: %s, min volume: %s", event.Currency, masterAmount.String(), currencyInstance.Min_volume.String())

		if masterAmount.LessThan(currencyInstance.Min_volume) {
			log.Infof("Liquidity low for %s. ", event.Currency)
			if conf.IsDevEnvironment() {
				time.Sleep(time.Duration(mockDelay[string(currencyInstance.Name)]) * time.Second)
			}

			depositAmount := currencyInstance.Max_volume.Sub(currencyInstance.Min_volume)
			providerAccount := models.FindPaymentAccountByAppEntityIdAndCurrencyName(models.GetProvider(), event.Currency)
			providerTigerId := providerAccount.TigerId

			transfers := []tigerTypes.Transfer{{
				ID:              tigerTypes.ID(),
				DebitAccountID:  tigerTypes.ToUint128(uint64(masterAccount.TigerId)),
				CreditAccountID: tigerTypes.ToUint128(uint64(providerTigerId)),
				Amount:          tigerTypes.ToUint128(depositAmount.BigInt().Uint64()),
				Ledger:          uint32(currencyInstance.Ledger),
				Code:            1,
				Flags:           0,
				Timestamp:       0,
			}}

			transferRes, err := tiger.Instance().CreateTransfers(transfers)
			if err != nil {
				log.Errorf("Failed to create transfer for %s: %v", event.Currency, err)
				continue
			}
			log.Infof("Transfer successful for %s. Transfer details: %+v", event.Currency, transferRes)
		}
	}
}

func SendPaymentEvent(currency models.CurrencyName) {
	event := PaymentEvent{
		Currency: currency,
	}

	select {
	case paymentChannel <- event:
		log.Infof("Payment event sent for currency: %s", currency)
	default:
		log.Error("Payment channel is full. Failed to send payment event.")
	}
}
