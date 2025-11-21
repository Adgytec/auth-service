package authentication

import "github.com/Adgytec/service-protos/auth/v1"

type authServicePC struct {
	service *authService
	auth.UnimplementedAuthServiceServer
}

func NewAuthServicePC() auth.AuthServiceServer {
	return &authServicePC{
		service: newAuthService(),
	}
}
