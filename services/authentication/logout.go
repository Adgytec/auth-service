package authentication

import (
	"context"
	"net/http"
)

func (s *authService) logout(ctx context.Context, sessionID string) error {
	return s.s.Delete(ctx, sessionID)
}

func (m *authServiceMux) logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie(cookieKeySessionID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	logoutErr := m.service.logout(r.Context(), session.Value)
	if logoutErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
