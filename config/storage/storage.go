package storage

import (
	"context"
)

type Storage interface {
	Get(ctx context.Context, sessionID string) (string, error)
	New(ctx context.Context, userID string) (string, error)
	Delete(ctx context.Context, sessionID string) error
}

type stubStorage struct{}

func New() Storage {
	return &stubStorage{}
}
