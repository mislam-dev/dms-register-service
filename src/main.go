package main

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
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

	file, _ := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	console := zerolog.ConsoleWriter{Out: os.Stderr}
	multi := zerolog.MultiLevelWriter(console, file)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	debug := flag.Bool("debug", false, "set log level to debug")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Info().Msg("Server starting on port 3002")

	if err := app.Listen(":3002"); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
