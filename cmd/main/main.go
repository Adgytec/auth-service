package main

import (
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Warn().
			Err(envErr).
			Msg("failed to load .env")
	}

	// add logger details
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	logLevel, parseErr := zerolog.ParseLevel(logLevelStr)
	if parseErr != nil {
		log.Warn().
			Err(parseErr).
			Str("log_level_provided", logLevelStr).
			Msg("invalid log level provided, defaulting to 'info'")
		logLevel = zerolog.InfoLevel // default
	}
	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var output io.Writer = os.Stderr
	if os.Getenv("ENV") == "development" {
		output = zerolog.ConsoleWriter{
			Out: os.Stderr,
			FieldsExclude: []string{
				zerolog.TimestampFieldName,
				"remote_ip",
				"user_agent",
			},
		}
	}
	log.Logger = log.Output(output)
}
