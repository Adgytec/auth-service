package authentication

import (
	"github.com/Adgytec/auth-service/utils/services"
	"github.com/go-chi/chi/v5"
)

type authServiceMux struct {
	service *authService
}

func (m *authServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	return mux
}

func NewServiceMux() services.Mux {
	return &authServiceMux{
		service: newAuthService(),
	}
}
