package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func (s *EkadashiServer) handleRegistration(username string) error {
	password := generateHash(username)
	userRequest, err := marshalMessage(username, password)
	if err != nil {
		return fmt.Errorf("cannot unmarshal user: %v", err)
	}
	resp, err := http.Post("http://localhost:9000/register", "application/json",
		userRequest)
	if err != nil {
		return fmt.Errorf("cannot send request: %v", err)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot registrate user: %v %v %v", resp.Status, resp.Header, string(response))
	}
	err = s.addUser(username, password)
	if err != nil{
		return fmt.Errorf("%v",err)
	}
	err = s.handleLogin(username)
	if err != nil{
		return fmt.Errorf("cannot login user: %v",err)
	}
	return nil
}

func (s *EkadashiServer) handleLogin(username string) error {
	password, err := s.getUserHash(username)
	if err != nil{
		return err
	}
	userRequst, err := marshalMessage(username,password )
	if err != nil {
		return fmt.Errorf("cannot marshal user: %v", err)
	}
	resp, err := http.Post("http://localhost:9000/login", "application/json", userRequst)
	if err != nil {
		return fmt.Errorf("cannot send request: %v", err)
	}
	err = s.setCookie(username, resp) //w8 until someone's would help me asap
	if err != nil{
		return fmt.Errorf("%v",err)
	}
	response,_ := ioutil.ReadAll(resp.Body) // it might be vice versa with previous err
	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot login user: %v %v %v", resp.Status, resp.Header, response)
	}
	return nil
}


