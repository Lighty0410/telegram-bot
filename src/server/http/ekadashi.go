package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// EkadashiDate contains information about ekadashi date.
type EkadashiDate struct {
	Date string `json:"date"`
}

const sessionName = "session_token"

// showEkadashiHandler shows next ekadashi day based on another server.
func (s *Service) ShowEkadashi(username string) (string, error) {
	user, err := s.controller.GetUser(username)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", s.serverURL+s.showEkadashiURL, nil)
	if err != nil {
		return "", fmt.Errorf("cannot get enpdoint: %v", err)
	}
	req.AddCookie(&http.Cookie{Name: sessionName, Value: user.Token})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("cannot send request: %v", err)
	}
	defer resp.Body.Close()
	ekadashi, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ekadashi), nil
}
