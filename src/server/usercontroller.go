package server

import (
	"fmt"
	"github.com/Lighty0410/telegram-bot/src/database"
)

func (s *EkadashiServer)addUser(username, password string)error{
	err := s.db.AddUser(&database.User{Username:username, Hash:password})
	if err != nil{
		return fmt.Errorf("cannot add user to database: %v",err)
	}
	return nil
}

func (s *EkadashiServer)getUserHash(username string)(string,error){
	hash, err := s.db.GetUser(username)
	if err != nil{
		return "", fmt.Errorf("cannot get user: %v",err)
	}
	return hash,nil
}
