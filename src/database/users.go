package database

import "fmt"

type User struct {
	Username string
	Token string
	Password string
}

func (s *Service)AddUser(u *User)error{
	s.users[u.Username] = u.Hash
	if len(s.users[u.Username]) == 0 {
		return fmt.Errorf("database cannot contain empty username or password")
	}
	return nil
}

func (s *Service)GetUser(username string)(string, error){
	if len(s.users[username]) == 0{
		return "", fmt.Errorf("this user doesn't exist in the database")
	}
	return s.users[username], nil
}
