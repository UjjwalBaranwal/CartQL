// Package logger provides a configured zerolog logger for the application.
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New creates and returns a new zerolog.Logger instance configured for the application.
func New() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	if os.Getenv("GIN_MODE") != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	return log.Logger
}
