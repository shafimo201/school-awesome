package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(level string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	out := zerolog.ConsoleWriter{Out: os.Stdout}
	logger := zerolog.New(out).With().Timestamp().Logger()

	switch level {
	case "debug":
		logger = logger.Level(zerolog.DebugLevel)
	case "info":
		logger = logger.Level(zerolog.InfoLevel)
	case "warn":
		logger = logger.Level(zerolog.WarnLevel)
	case "error":
		logger = logger.Level(zerolog.ErrorLevel)
	default:
		logger = logger.Level(zerolog.InfoLevel)
	}

	log.Logger = logger
	return logger
}

func Discard() zerolog.Logger {
	return zerolog.New(io.Discard).With().Timestamp().Logger()
}
