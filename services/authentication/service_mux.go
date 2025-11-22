package authentication

import (
	"errors"
	"fmt"
	"os"

	"github.com/Adgytec/auth-service/config/storage"
	"github.com/Adgytec/auth-service/utils/services"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/go-chi/chi/v5"
)

type authServiceMux struct {
	service    *authService
	jwtKeyfunc keyfunc.Keyfunc
}

func (m *authServiceMux) Router() *chi.Mux {
	mux := chi.NewMux()

	mux.Post("/new-session", m.newSession)
	mux.Post("/logout", m.logout)

	return mux
}

func NewServiceMux(s storage.Storage) (services.Mux, error) {
	userPoolID := os.Getenv("AWS_USER_POOL_ID")
	if userPoolID == "" {
		return nil, errors.New("can't find value for AWS_USER_POOL_ID env variable")
	}

	userPoolRegion := os.Getenv("AWS_USER_POOL_REGION")
	if userPoolRegion == "" {
		return nil, errors.New("can't find value for AWS_USER_POOL_REGION env variable")
	}

	jwkSetEndpoint := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", userPoolRegion, userPoolID)
	jwtKeyfunc, keyFuncErr := keyfunc.NewDefault([]string{jwkSetEndpoint})
	if keyFuncErr != nil {
		return nil, keyFuncErr
	}

	return &authServiceMux{
		service:    newAuthService(s),
		jwtKeyfunc: jwtKeyfunc,
	}, nil
}
