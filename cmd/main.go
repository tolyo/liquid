package main

import (
	"liquid/pkg/conf"
	"liquid/pkg/db"
	"liquid/pkg/rest/fxrate"
	"liquid/pkg/rest/master"
	paymentaccount "liquid/pkg/rest/payment_account"
	"liquid/pkg/rest/transfer"
	"liquid/pkg/services"
	"liquid/pkg/tiger"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envVarValue := os.Getenv("ENV")
	if envVarValue == "" {
		envVarValue = "DEV"
	}

	conf.LoadConfig((envVarValue))

	dbInstance := db.SetupInstance()
	defer dbInstance.Close()

	services.InitPaymentChannel()
	go services.ListenForPayments()

	tg := tiger.SetupInstance()
	defer tg.Close()

	services.InitProviderAccounts()
	services.InitMasterAccounts()
	services.InitTestAccount("USER_1")
	services.InitTestAccount("USER_2")

	app := fiber.New()
	app.Static("/", "./web")

	app.Get("/master", master.GetMasterBalances())
	app.Get("/balances/:externalid", paymentaccount.GetClientBalance())
	app.Post("/fx-rate", fxrate.UpdateRates())
	app.Post("/transfer", transfer.CreateTransfer())

	if err := app.Listen(":" + conf.Get().HttpPort); err != nil {
		// Log the error (you can also use log.Fatal or panic depending on your needs)
		log.Fatal("Failed to start server: " + err.Error())
	}
}
