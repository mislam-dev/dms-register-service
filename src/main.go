package main

import (
	database "mislam-dev/dms-register-service/src/core/config"
	"mislam-dev/dms-register-service/src/core/docs"
	"mislam-dev/dms-register-service/src/core/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "dms-register-service",
		ServerHeader: "dms",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server up and running",
		})
	})

	// logging setup
	file := logger.Setup(app)
	defer file.Close()

	// swaggger doc setup
	docs.SetUp(app)

	// database connection
	database.ConnectDatabase()
	defer database.CloseConnection()

	// routes setup

	if err := app.Listen(":3002"); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
