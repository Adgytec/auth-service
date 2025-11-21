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
	return &httpAndGRPCServer{}, nil
}
