package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const EkadashiToken = "EkadashiToken"

func ResponseEkadashiBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(EkadashiToken))
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
			err := registrateUser(strings.Split(update.Message.CommandArguments(), " "))
		case "login":

		}

	}
}

func createhash(str string) {
	fmt.Println(str)
}

var userMap = make(map[string]string)

func registrateUser(userInfo []string) error {
	message := map[string]interface{}{
		"username":"username",
		"password":"password",
	}
	userMap["user"] = "smt"
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:9000/register", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Status ==  {}
}

func sendRequest() string {
	resp, err := http.Get("http://localhost:9000/ekadashi/next")
	if err != nil {
		log.Fatalln(err)
	}
	smt,_ := ioutil.ReadAll(resp.Body)
	return string(smt)
}
