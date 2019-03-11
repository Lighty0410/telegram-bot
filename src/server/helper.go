package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Lighty0410/telegram-bot/telegram-bot/src/database"
	"net/http"
)

type message struct {
	Username string
	Password string
}

func marshalMessage(username, hash string) (*bytes.Buffer, error) {
	userMessage := message{Username: username, Password: hash}
	jsonMessage, err := json.Marshal(userMessage)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal user: %v", err)
	}
	userRequest := bytes.NewBuffer(jsonMessage)
	return userRequest, nil
}

func (s *EkadashiServer) getCookieValue(username string) (string, error) {
	cookieValue, err := s.db.GetCookie(username)
	if err != nil {
		return "", fmt.Errorf("cannot get cookie: %v", err)
	}
	return cookieValue, nil
}

func (s *EkadashiServer) setCookie(username string, cookie []*http.Cookie) error {
	var cookieValue string
	for _, cookie := range cookie {
		if cookie.Name == "" || cookie.Value == "" {
			return fmt.Errorf("this cookie doesn't exist")
		}
		cookieValue = cookie.Value
	}
	err := s.db.SetCookie(username, &database.User{Token: cookieValue})
	if err != nil {
		return err
	}
	return nil
}
