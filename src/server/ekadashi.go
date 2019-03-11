package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EkadashiDate struct {
	Date string `json:"date"`
}

const sessionName = "session_token"

func (s *EkadashiBot) showEkadashiHandler(username string) (string, error) {
	user, err := s.getUser(username)
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
	ekadashi, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var ekadashiDate EkadashiDate
	err = json.Unmarshal(ekadashi, &ekadashiDate)
	if err != nil {
		return "", err
	}
	return ekadashiDate.Date, nil
}
