package memdb

import "fmt"

var ErrUserNotFound = fmt.Errorf("user not found")

// UpsertUser add and updates user in the memory database.
func (s *Service) UpsertUser(u User) error {
	s.users[u.ID] = u
	return nil
}

// GetUser retrieves information from the memory database and sends it to another function.
func (s *Service) GetUser(userID string) (User, error) {
	u, ok := s.users[userID]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
