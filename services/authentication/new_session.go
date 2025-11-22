package authentication

import (
	"context"
	"net/http"
	"time"
)

const (
	authorizationHeaderKey    = "Authorization"
	authorizationSchemeBearer = "bearer"
)

var sessionCookieDuration = 24 * 15 * time.Hour

func (s *authService) newSession(ctx context.Context, userID string) (string, error) {
	return s.s.New(ctx, userID)
}

func (m *authServiceMux) newSession(w http.ResponseWriter, r *http.Request) {
	userID, err := m.validateAndGetUserIDCtx(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionID, err := m.service.newSession(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     cookieKeySessionID,
		Value:    sessionID,
		Path:     "/",
		Domain:   m.cookieDomain,
		Expires:  time.Now().Add(sessionCookieDuration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusCreated)
}
