package authentication

import (
	"github.com/Adgytec/auth-service/config/storage"
	"github.com/Adgytec/service-protos/auth/v1"
)

type authServicePC struct {
	service *authService
	auth.UnimplementedAuthServiceServer
}

func NewAuthServicePC(s storage.Storage) auth.AuthServiceServer {
	return &authServicePC{
		service: newAuthService(s),
	}
}
