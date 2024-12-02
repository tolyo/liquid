package services

import (
	"liquid/pkg/models"

	"github.com/shopspring/decimal"
)

var FxRates = map[models.CurrencyPair]decimal.Decimal{
	"USD/EUR": decimal.NewFromFloat(0.9217),
	"EUR/USD": decimal.NewFromFloat(1.085),
	"USD/JPY": decimal.NewFromFloat(110.25),
	"JPY/USD": decimal.NewFromFloat(0.0091),
	"USD/GBP": decimal.NewFromFloat(0.75),
	"GBP/USD": decimal.NewFromFloat(1.3333),
	"USD/AUD": decimal.NewFromFloat(1.35),
	"AUD/USD": decimal.NewFromFloat(0.7407),
	"EUR/JPY": decimal.NewFromFloat(129.53),
	"JPY/EUR": decimal.NewFromFloat(0.0077),
	"EUR/GBP": decimal.NewFromFloat(0.85),
	"GBP/EUR": decimal.NewFromFloat(1.1765),
	"EUR/AUD": decimal.NewFromFloat(1.6),
	"AUD/EUR": decimal.NewFromFloat(0.625),
	"GBP/JPY": decimal.NewFromFloat(150.45),
	"JPY/GBP": decimal.NewFromFloat(0.0066),
	"GBP/AUD": decimal.NewFromFloat(1.8),
	"AUD/GBP": decimal.NewFromFloat(0.5556),
	"AUD/JPY": decimal.NewFromFloat(82.5),
	"JPY/AUD": decimal.NewFromFloat(0.0121),
}
