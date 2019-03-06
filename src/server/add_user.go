package server

import (
	"bytes"
	"fmt"
	"net/http"
)

func loginUser(username string, hash string) error{
	userRequst, err := marshalUsers(username, hash)
	if err != nil{
		return fmt.Errorf("cannot marshal user: %v",err)
	}
	resp, err := http.Post("http://localhost:9000/login", "application/json", bytes.NewBuffer(userRequst))
	if err != nil{
		return fmt.Errorf("cannot send request: %v", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot login user: %v %v %v", resp.Status, resp.Header, resp.Body)
	}
	addUsers(resp, username)
	err = doesUserExist(username)
	if err != nil{
		return fmt.Errorf("%v",err)
	}
	return nil
}

func registerUser(username string, password string) error {
	userRequest, err := marshalUsers(username, password)
	if err != nil{
		return fmt.Errorf("cannot unmarshal user: %v", err)
	}
	resp, err := http.Post("http://localhost:9000/register", "application/json",
		bytes.NewBuffer(userRequest))
	if err != nil {
		return fmt.Errorf("cannot send request: %v",err)
	}
	if resp.StatusCode != 200{
		return fmt.Errorf("cannot registrate user: %v %v %v", resp.Status, resp.Header, resp.Body)
	}
	err = loginUser(username, password)
	if err != nil{
		return fmt.Errorf("an error occured while login user: %v",err)
	}
	return nil
}
