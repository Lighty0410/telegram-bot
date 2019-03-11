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

func (s *EkadashiServer) showEkadashiHandler(username string) (string, error) {
	token, err := s.getCookieValue(username)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", "http://localhost:9000/ekadashi/next", nil)
	if err != nil {
		return "", fmt.Errorf("cannot get enpdoint: %v", err)
	}
	req.AddCookie(&http.Cookie{Name: "session_token", Value: token}) // temporary name
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("cannot send request: %v", err)
	}
	ekadashi, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ekadashiDate := EkadashiDate{}
	err = json.Unmarshal(ekadashi, &ekadashiDate.Date)
	if err != nil {
		return "", err
	}
	return ekadashiDate.Date, nil
}
