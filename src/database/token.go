package database

import "fmt"

type Token struct {
	Name string
	Hash string
}
func (s *Service)SetCookie(t *Token)error{
	if len(s.users[t.Name]) == 0{
		return fmt.Errorf("field users cannot be empty")
	}
	s.users[t.Name] = t.Hash
	return nil
}

func (s *Service) GetCookie(t *Token)(string,error){
	if len(s.users[t.Name]) == 0{
		return "", fmt.Errorf("field users cannot be empty")
	}
	return s.users[t.Name], nil
}
