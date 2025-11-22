package router

import (
	"os"
	"strings"

	"github.com/Adgytec/auth-service/services/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

func NewHTTPRouter() (*chi.Mux, error) {
	log.Info().Msg("adding application mux")
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.Heartbeat("/health"))

	originEnv := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := []string{}

	if originEnv != "" {
		allowedOrigins = strings.Split(originEnv, ",")
	} else {
		log.Warn().Msg("ALLOWED_ORIGINS not set")
		if os.Getenv("ENV") == "development" {
			allowedOrigins = []string{"http://localhost:*"}
		}
	}

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	authMux, authMuxErr := authentication.NewServiceMux()
	if authMuxErr != nil {
		return nil, authMuxErr
	}

	mux.Mount("/", authMux.Router())

	return mux, nil
}
