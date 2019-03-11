package server

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type message struct {
	Username string
	Password string
}

func marshalMessage(username, password string) (*bytes.Buffer, error) {
	userMessage := message{Username: username, Password: password}
	jsonMessage, err := json.Marshal(userMessage)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal user: %v", err)
	}
	userRequest := bytes.NewBuffer(jsonMessage)
	return userRequest, nil
}
