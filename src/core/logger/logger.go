package logger

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(app *fiber.App) *os.File {
	file, _ := os.OpenFile("./logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

	return file
}
