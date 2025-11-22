package storage

import (
	"context"
	"errors"
)

type Storage interface {
	Get(ctx context.Context, sessionID string) (string, error)
	New(ctx context.Context, userID string) (string, error)
	Delete(ctx context.Context, sessionID string) error
}

type stubStorage struct{}

func (s *stubStorage) Get(ctx context.Context, sessionID string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *stubStorage) New(ctx context.Context, userID string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *stubStorage) Delete(ctx context.Context, sessionID string) error {
	return errors.New("not implemented")
}

func New() Storage {
	return &stubStorage{}
}
