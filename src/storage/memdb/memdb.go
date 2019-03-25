package memdb

import "github.com/Lighty0410/telegram-bot/src/storage"

// Service struct contains a simple RAM-based database in map-like format.
type Service struct {
	users map[string]storage.User
}

// NewService creates database in memory database.
func NewService() *Service {
	return &Service{
		users: make(map[string]storage.User),
	}
}
