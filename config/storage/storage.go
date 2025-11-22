package storage

import "errors"

type Storage interface {
	Get(sessionID string) (string, error)
	New(userID string) (string, error)
	Delete(sessionID string) error
}

type stubStorage struct{}

func (s *stubStorage) Get(sessionID string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *stubStorage) New(userID string) (string, error) {
	return "", errors.New("not implemented")
}

func (s *stubStorage) Delete(sessionID string) error {
	return errors.New("not implemented")
}

func New() Storage {
	return &stubStorage{}
}
