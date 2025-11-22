package authentication

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Adgytec/auth-service/utils/core"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type authContextKey string

const (
	authContextKeyUserID authContextKey = "userID"
)

// validateAccessToken() parses and validate jwt token and return user id
func (m *authServiceMux) validateAccessToken(accessToken string) (uuid.UUID, error) {
	jwtToken, jwtParseErr := jwt.Parse(accessToken, m.jwtKeyfunc.Keyfunc)
	if jwtParseErr != nil {
		if errors.Is(jwtParseErr, jwt.ErrTokenMalformed) ||
			errors.Is(jwtParseErr, jwt.ErrTokenSignatureInvalid) ||
			errors.Is(jwtParseErr, jwt.ErrTokenExpired) ||
			errors.Is(jwtParseErr, jwt.ErrTokenNotValidYet) {
			return uuid.Nil, errors.New("invalid access token")
		}

		return uuid.Nil, jwtParseErr
	}

	invalidTokenErr := errors.New("invalid token")

	if !jwtToken.Valid {
		return uuid.Nil, invalidTokenErr
	}

	// get username from claims
	claims, claimsOK := jwtToken.Claims.(jwt.MapClaims)
	if !claimsOK {
		return uuid.Nil, invalidTokenErr
	}

	// cognito access token contains claim 'username' for user's username field
	username, usernameOK := claims["username"].(string)
	if !usernameOK {
		return uuid.Nil, invalidTokenErr
	}

	userID := core.GetUserIDFromUsername(username)
	return userID, nil
}

// return scheme, value, error
func parseAuthHeader(authHeader string) (string, string, error) {
	authHeaderSlice := strings.Fields(authHeader)
	if len(authHeaderSlice) != 2 {
		return "", "", errors.New("invalid auth header")
	}

	scheme := strings.ToLower(authHeaderSlice[0])
	value := authHeaderSlice[1]
	return scheme, value, nil
}

func (m *authServiceMux) validateAndGetUserID(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorizationHeaderKey)
	scheme, value, authErr := parseAuthHeader(authHeader)
	if authErr != nil {
		return "", authErr
	}

	switch scheme {
	case authorizationSchemeBearer:
		userID, validateErr := m.validateAccessToken(value)
		if validateErr != nil {
			return "", validateErr
		}
		return userID.String(), nil

	default:
		return "", errors.New("unsupported auth scheme")
	}
}
