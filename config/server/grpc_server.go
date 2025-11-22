package server

import (
	"errors"
	"net"
	"os"

	"github.com/Adgytec/auth-service/services/authentication"
	"github.com/Adgytec/service-protos/auth/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const defaultGRPCPort = "50051"

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
	port     string
}

func (s *grpcServer) ListenAndServe() error {
	log.Info().
		Str("port", s.port).
		Msg("grpc server started listening")

	return s.server.Serve(s.listener)
}

func (s *grpcServer) Shutdown() error {
	log.Info().Msg("grpc server shutting down")

	s.server.GracefulStop()
	listenerErr := s.listener.Close()
	if !errors.Is(listenerErr, net.ErrClosed) {
		return listenerErr
	}

	return nil
}

func newGRPCServer() (Server, error) {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = defaultGRPCPort
		log.Warn().
			Msgf("missing GRPC_PORT env variable, using default grpc port: %s", defaultGRPCPort)
	}

	listener, listenerErr := net.Listen("tcp", ":"+grpcPort)
	if listenerErr != nil {
		return nil, listenerErr
	}

	serverRegistrar := grpc.NewServer()
	auth.RegisterAuthServiceServer(serverRegistrar, authentication.NewAuthServicePC())

	return &grpcServer{
		server:   serverRegistrar,
		listener: listener,
		port:     grpcPort,
	}, nil
}
