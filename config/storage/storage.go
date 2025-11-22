package storage

type Storage interface {
	Get(sessionID string) (string, error)
	New(userID string) (string, error)
	Delete(sessionID string) error
}
