package router

import (
	"github.com/Adgytec/auth-service/services/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

func NewHTTPRouter() *chi.Mux {
	log.Info().Msg("adding application mux")
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.Heartbeat("/health"))

	allowedOrigins := []string{
		"https://accounts.adgytec.in",
	}

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Mount("/", authentication.NewServiceMux().Router())

	return mux
}
