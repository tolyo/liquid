package paymentaccount

import (
	"liquid/pkg/models"
	"liquid/pkg/rest/mappers"
	"liquid/pkg/services"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Updated to return an error, which is required by fiber v2
func GetClientBalance() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info("/balances " + c.Params("externalid"))
		fetched := services.GetFullPaymentAccounts(models.AppEntityExternalId(c.Params("externalid")))
		return c.JSON(mappers.PaymentAccoutResponse(fetched)) // Return the JSON response
	}
}
