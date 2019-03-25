package http

import "fmt"

package server

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
)

type errorMessage struct {
	Reason string `json:"reason"`
}

// handleRegistration register user in the microservice.
func (s *EkadashiBot) handleRegistration(username string) error {
	password := generateHash(username)
	userRequest, err := marshalMessage(username, password)
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
	err = s.addUser(User{ID: username, Password: password})
	if err != nil {
		return err
	}
	err = s.handleLogin(username)
	if err != nil {
		return fmt.Errorf("cannot login user: %v", err)
	}
	return nil
}

// handleLogin login user in the microservice.
func (s *EkadashiBot) handleLogin(username string) error {
	user, err := s.getUser(username)
	if err != nil {
		return err
	}
	userRequest, err := marshalMessage(username, user.Password)
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
	err = s.addUser(User{ID: user.ID, Password: user.Password, Token: cookieValue})
	if err != nil {
		return err
	}
	return nil
}

