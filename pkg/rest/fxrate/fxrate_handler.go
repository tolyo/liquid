package fxrate

import (
	"liquid/pkg/models"
	"liquid/pkg/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type FXrate struct {
	Pair      string    `json:"pair"`
	Rate      string    `json:"rate"`
	Timestamp time.Time `json:"timestamp"`
}

func UpdateRates() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tick FXrate

		if err := c.BodyParser(&tick); err != nil {
			log.Error(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		log.Infof("Received Tick: %s, Rate: %s, Timestamp: %s\n", tick.Pair, tick.Rate, tick.Timestamp)
		res, err := decimal.NewFromString(tick.Rate)
		if err != nil {
			log.Error(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}
		services.FxRates[models.CurrencyPair(tick.Pair)] = res

		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}
