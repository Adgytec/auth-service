package authentication

import (
	"context"
	"errors"

	"github.com/Adgytec/service-protos/auth/v1"
)

func (s *authService) getSessionInfo(ctx context.Context, sessionID string) (string, error) {
	return s.s.Get(ctx, sessionID)
}

func (a *authServicePC) GetSessionInfo(ctx context.Context, req *auth.GetSessionInfoRequest) (*auth.GetSessionInfoResponse, error) {
	if req == nil {
		return nil, errors.New("session id not present")
	}

	userID, sessionErr := a.service.getSessionInfo(ctx, req.SessionId)
	if sessionErr != nil {
		return nil, sessionErr
	}

	sessionValid := userID != ""
	return &auth.GetSessionInfoResponse{
		Valid:  sessionValid,
		UserId: userID,
	}, nil
}
