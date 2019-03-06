package server

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var userMap = make(map[string]string)

const EkadashiToken = "EkadashiToken"

func ResponseEkadashiBot() {
	bot, err := tgbotapi.NewBotAPI("705163703:AAEIbWGS_nobktv1URXi4_PXyjr7DfHvyhk")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates{
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "register":
			err := registerUser(strconv.FormatInt(int64(update.Message.Chat.ID),10), update.Message.CommandArguments())
			if err != nil {
				log.Println("cannot register user: ",err)
			}
		case "password":
			err := loginUser(strconv.FormatInt(int64(update.Message.Chat.ID),10), update.Message.CommandArguments())
			if err != nil {
				log.Println("cannot login user: ",err)
			}
		case "showEkadashi":
			ekadashiDate, err := showEkadashiHandler(strconv.FormatInt(int64(update.Message.Chat.ID),10))
			if err != nil{
				log.Println(err)
			}
			update.Message.Text = ekadashiDate
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,update.Message.Text)
			bot.Send(msg)
		}
	}
}


func showEkadashiHandler(password string)(string, error) {
	token := userMap[password]
	req, err := http.NewRequest("GET","http://localhost:9000/ekadashi/next", nil)
	if err != nil {
		return "", fmt.Errorf("cannot get enpdoint: %v",err)
	}
	req.AddCookie(&http.Cookie{Name:"session_token",Value: token}) // temporary name
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return "", fmt.Errorf("cannot send request: %v",err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("invalid request: %v %v %v", resp.StatusCode, resp.Header, resp.Body)
	}
	ekadashi,_ := ioutil.ReadAll(resp.Body)
	return string(ekadashi), nil
}
