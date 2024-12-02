package master

import (
	"liquid/pkg/rest/mappers"
	"liquid/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func GetMasterBalances() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched := services.GetMasterAccounts()
		return c.JSON(mappers.PaymentAccoutResponse(fetched))
	}
}
