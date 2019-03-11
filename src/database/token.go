package database

import "fmt"

func (s *Service) SetCookie(userID string, u *User) error {
	if len(s.users[userID].Password) == 0 {
		return fmt.Errorf("field users cannot be empty")
	}
	token := s.users[userID]
	token.Token = u.Token
	s.users[userID] = token
	return nil
}

func (s *Service) GetCookie(userID string) (string, error) {
	if len(s.users[userID].Password) == 0 {
		return "", fmt.Errorf("field users cannot be empty")
	}
	return s.users[userID].Token, nil
}
