package authentication

import (
	"github.com/Adgytec/service-protos/auth/v1"
)

type authService struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthService() auth.AuthServiceServer {
	return &authService{}
}
