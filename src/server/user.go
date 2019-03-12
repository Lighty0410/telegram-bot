package server

import (
	"fmt"

	"github.com/Lighty0410/telegram-bot/src/memdb"
)

type User struct {
	ID       string
	Token    string
	Password string
}

// ErrNoUser is an error that indicates that there's no such user in microservice.
var ErrNoUser = fmt.Errorf("user not found")

func (s *EkadashiBot) addUser(u User) error {
	err := s.db.UpsertUser(memdb.User{ID: u.ID, Password: u.Password, Token: u.Token})
	if err != nil {
		return fmt.Errorf("cannot add user to memdb: %v", err)
	}
	return nil
}

func (s *EkadashiBot) getUser(username string) (User, error) {
	user, err := s.db.GetUser(username)
	if err == memdb.ErrUserNotFound {
		return User{}, ErrNoUser
	}
	if err != nil {
		return User{}, fmt.Errorf("cannot get user: %v", err)
	}
	return User{
		ID:       user.ID,
		Token:    user.Token,
		Password: user.Password,
	}, nil
}
