package memdb

import (
	"github.com/Lighty0410/telegram-bot/src/storage"
)

// UpsertUser adds and updates user.
func (s *Service) UpsertUser(u storage.User) error {
	s.users[u.ID] = u
	return nil
}

// GetUser retrieves information from the memory database and returns user or error.
func (s *Service) GetUser(userID string) (storage.User, error) {
	u, ok := s.users[userID]
	if !ok {
		return storage.User{}, storage.ErrUserNotFound
	}
	return u, nil
}
