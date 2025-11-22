package core

import (
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var (
	idNamespace uuid.UUID
	idOnce      sync.Once
)

// this method panics if id namespace is not found
func getIDNamespace() uuid.UUID {
	idOnce.Do(func() {
		namespaceString := os.Getenv("ID_NAMESPACE")
		namespaceVal, namespaceErr := uuid.Parse(namespaceString)
		if namespaceErr != nil {
			log.Fatal().
				Err(namespaceErr).
				Str("action", "get id namespace").
				Send()
		}
		idNamespace = namespaceVal
	})

	return idNamespace
}

func GetIDFromPayload(payload []byte) uuid.UUID {
	namespace := getIDNamespace()
	return uuid.NewSHA1(namespace, payload)
}

func GetUserIDFromUsername(username string) uuid.UUID {
	return GetIDFromPayload([]byte(strings.ToLower(strings.TrimSpace(username))))
}
