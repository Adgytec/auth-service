package storage

import (
	"context"
	"errors"
)

func (s *stubStorage) Delete(ctx context.Context, sessionID string) error {
	return errors.New("not implemented")
}
