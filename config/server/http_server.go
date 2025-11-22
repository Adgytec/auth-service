package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Adgytec/auth-service/config/router"
	"github.com/Adgytec/auth-service/config/storage"
	"github.com/rs/zerolog/log"
)

const defaultHTTPPort = "8080"

type httpServer struct {
	server *http.Server
	port   string
}

func (s *httpServer) ListenAndServe() error {
	log.Info().
		Str("port", s.port).
		Msg("http server started listening")

	return s.server.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	log.Info().Msg("http server shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return s.server.Shutdown(shutdownCtx)
}

func newHTTPServer(s storage.Storage) (Server, error) {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = defaultHTTPPort
		log.Warn().
			Msgf("missing HTTP_PORT env variable, using default http port: %s", defaultHTTPPort)
	}

	mux, muxErr := router.NewHTTPRouter(s)
	if muxErr != nil {
		return nil, muxErr
	}

	var protocols http.Protocols
	protocols.SetUnencryptedHTTP2(true)

	appServer := http.Server{
		Addr:              ":" + httpPort,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           mux,
		Protocols:         &protocols,
	}

	return &httpServer{
		server: &appServer,
		port:   httpPort,
	}, nil
}
