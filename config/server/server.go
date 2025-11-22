package server

import (
	"errors"

	"github.com/Adgytec/auth-service/config/storage"
)

type Server interface {
	ListenAndServe() error
	Shutdown() error
}

type httpAndGRPCServer struct {
	httpServer Server
	grpcServer Server
}

func (s *httpAndGRPCServer) ListenAndServe() error {
	errCh := make(chan error, 2)

	// Start HTTP server
	go func() {
		errCh <- s.httpServer.ListenAndServe()
	}()

	// Start gRPC server
	go func() {
		errCh <- s.grpcServer.ListenAndServe()
	}()

	return <-errCh
}

func (s *httpAndGRPCServer) Shutdown() error {
	httpShutdownErr := s.httpServer.Shutdown()
	grpcShutdownErr := s.grpcServer.Shutdown()
	return errors.Join(httpShutdownErr, grpcShutdownErr)
}

func NewServer() (Server, error) {
	// create new storage
	s := storage.New()

	httpServer, httpServerErr := newHTTPServer(s)
	if httpServerErr != nil {
		return nil, httpServerErr
	}

	grpcServer, grpcServerErr := newGRPCServer(s)
	if grpcServerErr != nil {
		return nil, grpcServerErr
	}

	return &httpAndGRPCServer{
		httpServer: httpServer,
		grpcServer: grpcServer,
	}, nil
}
