package storage

import (
	"context"
	"errors"
)

func (s *stubStorage) Get(ctx context.Context, sessionID string) (string, error) {
	return "", errors.New("not implemented")
}
