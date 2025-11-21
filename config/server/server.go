package server

type Server interface {
	ListenAndServe() error
	Shutdown() error
}

type httpAndGRPCServer struct {
	httpServer Server
	grpcServer Server
}

func (s *httpAndGRPCServer) ListenAndServe() error {
	httpServerErr := s.httpServer.ListenAndServe()
	if httpServerErr != nil {
		return httpServerErr
	}

	return s.grpcServer.ListenAndServe()
}

func (s *httpAndGRPCServer) Shutdown() error {
	httpShutdownErr := s.httpServer.Shutdown()
	grpcShutdownErr := s.grpcServer.Shutdown()
	if httpShutdownErr != nil {
		return httpShutdownErr
	}

	return grpcShutdownErr
}

func NewServer() (Server, error) {
	httpServer, httpServerErr := newHTTPServer()
	if httpServerErr != nil {
		return nil, httpServerErr
	}

	grpcServer, grpcServerErr := newGRPCServer()
	if grpcServerErr != nil {
		return nil, grpcServerErr
	}

	return &httpAndGRPCServer{
		httpServer: httpServer,
		grpcServer: grpcServer,
	}, nil
}
