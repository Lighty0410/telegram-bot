package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func marshalUsers(username, hash string)([]byte,error){
	message := map[string]interface{}{
		"username": username,
		"password": hash,
	}
	log.Println(username)
	userRequest, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal user: %v", err)
	}
	return userRequest, nil
}

func doesUserExist(userInfo string)error{
	token := userMap[userInfo]
	if token == ""{
		return fmt.Errorf("user doesn't exist")
	}
	return nil
}

func addUsers(token *http.Response, userInfo string){
	var cookieValue string
	var cookieName string
	for _, hash := range token.Cookies(){
		cookieName = hash.Name
		cookieValue = hash.Value
	}
	userMap[userInfo] = cookieValue
	userMap[cookieName] = "session_token" // temporary
}
