package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Adgytec/auth-service/config/server"
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
	if parseErr != nil || logLevel == zerolog.NoLevel {
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

	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	appServer, serverErr := server.NewServer()
	if serverErr != nil {
		log.Fatal().
			Err(serverErr).
			Str("action", "new server").
			Send()
	}

	go func() {
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("server error, triggering shutdown")
			stop()
		}
	}()

	<-rootCtx.Done()

	// gracefull shutdown for server here
	if err := appServer.Shutdown(); err != nil {
		log.Error().
			Err(err).
			Str("action", "server shutdown").
			Send()
	} else {
		log.Info().
			Msg("server shutdown gracefully")
	}
}
