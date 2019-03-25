package controller

import (
	"fmt"

	"github.com/Lighty0410/telegram-bot/src/storage"
)

// User is a struct that contains user's info.
type User struct {
	ID       string
	Token    string
	Password string
}

// Controller is an object that provides an access for the controller's functionality.
type Controller struct {
	service storage.Service
}

// CreateController creates a new instance for the controller.
func NewController(service storage.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// ErrNoUser is an error that indicates that there's no such user in microservice.
var ErrNoUser = fmt.Errorf("user not found")

// AddUser adds user to the database.
// If succeed returns nil.
func (s *Controller) AddUser(u User) error {
	err := s.service.UpsertUser(storage.User{ID: u.ID, Password: u.Password, Token: u.Token})
	if err != nil {
		return fmt.Errorf("cannot add user to memdb: %v", err)
	}
	return nil
}

// GetUser gets user from the database.
// If succeed returns User structure and nil.
func (s *Controller) GetUser(username string) (User, error) {
	user, err := s.service.GetUser(username)
	if err == storage.ErrUserNotFound {
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
