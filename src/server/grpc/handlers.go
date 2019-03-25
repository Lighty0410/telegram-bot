package grpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/Lighty0410/telegram-bot/src/crypto"

	"github.com/Lighty0410/telegram-bot/src/server/controller"
	api "github.com/Lighty0410/telegram-bot/src/server/grpc/api"
)

// HandleRegistration register user in the microservice.
func (s *Service) Register(username string) error {
	password := crypto.GenerateHash(username)
	_, err := s.client.Register(context.Background(), &api.RegisterRequest{
		User: &api.User{
			Name:     username,
			Password: password,
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			err = s.controller.AddUser(controller.User{ID: username, Password: password})
			if err != nil {
				return fmt.Errorf("cannot add user to the database: %v", err)
			}
			err := s.HandleLogin(username)
			if err != nil {
				return fmt.Errorf("cannot login user: %v", err)
			}
			return nil
		}
		return fmt.Errorf("cannot register user: %v", err)
	}
	err = s.controller.AddUser(controller.User{ID: username, Password: password})
	if err != nil {
		return fmt.Errorf("cannot get add user to the database: %v", err)
	}
	err = s.HandleLogin(username)
	if err != nil {
		return fmt.Errorf("cannot login user: %v", err)
	}
	return nil
}

// HandleLogin login user in the microservice.
func (s *Service) HandleLogin(username string) error {
	user, err := s.controller.GetUser(username)
	if err != nil {
		return fmt.Errorf("cannot get user from the database: %v", err)
	}
	response, err := s.client.Login(context.Background(), &api.LoginRequest{
		User: &api.User{
			Name:     username,
			Password: user.Password,
		},
	})
	if err != nil {
		return fmt.Errorf("cannot get response from gRPC server: %v", err)
	}
	err = s.controller.AddUser(controller.User{ID: user.ID, Password: user.Password, Token: response.Auth.Token})
	if err != nil {
		return err
	}
	return nil
}
