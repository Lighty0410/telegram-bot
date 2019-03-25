package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Lighty0410/telegram-bot/src/crypto"
	"github.com/Lighty0410/telegram-bot/src/helper"
	"github.com/Lighty0410/telegram-bot/src/server/controller"
)

type errorMessage struct {
	Reason string `json:"reason"`
}

// Register registers user in the microservice.
func (s *HttpService) Register(username string) error {
	password := crypto.GenerateHash(username)
	userRequest, err := helper.MarshalMessage(username, password)
	if err != nil {
		return fmt.Errorf("cannot unmarshal user: %v", err)
	}
	resp, err := http.Post(s.serverURL+s.registerURL, "application/json", userRequest)
	if err != nil {
		return fmt.Errorf("cannot send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errorMessage := new(errorMessage)
		err := json.NewDecoder(resp.Body).Decode(&errorMessage)
		if err != nil {
			return fmt.Errorf("cannot decode request: %v", err)
		}
		if resp.StatusCode != http.StatusConflict {
			return fmt.Errorf("cannot register user: %v %v %v", resp.Status, resp.Header, errorMessage.Reason)
		}
	}
	err = s.controller.AddUser(controller.User{ID: username, Password: password})
	if err != nil {
		return err
	}
	err = s.Login(username)
	if err != nil {
		return fmt.Errorf("cannot login user: %v", err)
	}
	return nil
}

// Login logins user in the microservice.
func (s *HttpService) Login(username string) error {
	user, err := s.controller.GetUser(username)
	if err != nil {
		return err
	}
	userRequest, err := helper.MarshalMessage(username, user.Password)
	if err != nil {
		return fmt.Errorf("cannot marshal user: %v", err)
	}
	resp, err := http.Post(s.serverURL+s.loginURL, "application/json", userRequest)
	if err != nil {
		return fmt.Errorf("cannot send request: %v", err)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot login user: %v %v %v", resp.Status, resp.Header, response)
	}
	var cookieValue string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == sessionName {
			cookieValue = cookie.Value
			break
		}
	}
	err = s.controller.AddUser(controller.User{ID: user.ID, Password: user.Password, Token: cookieValue})
	if err != nil {
		return err
	}
	return nil
}
