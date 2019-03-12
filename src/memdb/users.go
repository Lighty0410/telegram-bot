package memdb

import "fmt"

// ErrUserNotFound returns an error when database cannot find the user.
var ErrUserNotFound = fmt.Errorf("user not found")

// UpsertUser adds and updates user.
func (s *Service) UpsertUser(u User) error {
	s.users[u.ID] = u
	return nil
}

// GetUser retrieves information from the memory database and returns user or error.
func (s *Service) GetUser(userID string) (User, error) {
	u, ok := s.users[userID]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
