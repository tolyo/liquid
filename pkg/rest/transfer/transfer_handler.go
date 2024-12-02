package transfer

import (
	"liquid/pkg/models"
	"liquid/pkg/services"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	tigerTypes "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

type PaymentRequest struct {
	UserId       string `json:"externalId"`
	Amount       int    `json:"amount"`
	CurrencySell string `json:"currencySell"`
	CurrencyBuy  string `json:"currencyBuy"`
}

func CreateTransfer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info("CreateTransfer Handler ")
		var payment PaymentRequest
		// Parse the JSON request body into the newProduct struct
		if err := c.BodyParser(&payment); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		log.Infof("Data %+v", payment)
		res, err := services.CreateExchange(
			models.AppEntityExternalId(payment.UserId),
			payment.Amount,
			models.CurrencyName(payment.CurrencySell),
			models.CurrencyName(payment.CurrencyBuy),
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(nil)
		}

		if res != tigerTypes.TransferOK {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"data": res.String(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data": res.String(),
		})
	}
}
