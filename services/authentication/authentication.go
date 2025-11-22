package authentication

import "github.com/Adgytec/auth-service/config/storage"

type authService struct {
	s storage.Storage
}

func newAuthService(s storage.Storage) *authService {
	return &authService{
		s: s,
	}
}
