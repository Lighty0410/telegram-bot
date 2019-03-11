package database

import "fmt"

func (s *Service) AddUser(userID string, u *User) error {
	s.users[userID] = User{Password: u.Password}
	if len(s.users[userID].Password) == 0 {
		return fmt.Errorf("database cannot contain empty username or password")
	}
	return nil
}

func (s *Service) GetUser(userID string) (string, error) {
	if len(s.users[userID].Password) == 0 {
		return "", fmt.Errorf("this user doesn't exist in the database")
	}
	return s.users[userID].Password, nil
}
