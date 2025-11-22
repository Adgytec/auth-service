package storage

import (
	"context"
	"errors"
)

func (s *stubStorage) New(ctx context.Context, userID string) (string, error) {
	return "", errors.New("not implemented")
}
