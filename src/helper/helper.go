package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type message struct {
	Username string
	Password string
}

// MarshalMessage is a helper to marshal message from string to *bytes.Buffer.
// It's necessary because we can use message in requests properly.
func MarshalMessage(username, password string) (*bytes.Buffer, error) {
	userMessage := message{Username: username, Password: password}
	jsonMessage, err := json.Marshal(userMessage)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal user: %v", err)
	}
	userRequest := bytes.NewBuffer(jsonMessage)
	return userRequest, nil
}
