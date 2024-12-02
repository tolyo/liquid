package mappers

import (
	"liquid/pkg/models"

	"github.com/gofiber/fiber"
)

type PaymentAccount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func PaymentAccoutResponse(data []models.FullPaymentAccount) *fiber.Map {
	return &fiber.Map{
		"data": data,
	}
}
