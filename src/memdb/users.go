package memdb

import "fmt"

var ErrUserNotFound = fmt.Errorf("user not found")

func (s *Service) UpsertUser(u User) error {
	s.users[u.ID] = u
	return nil
}

func (s *Service) GetUser(userID string) (User, error) {
	u, ok := s.users[userID]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return u, nil
}
