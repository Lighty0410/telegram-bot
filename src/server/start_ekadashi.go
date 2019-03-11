package server

import (
	"github.com/Lighty0410/telegram-bot/telegram-bot/src/database"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

const ekadashiToken = "EKADASHI_TOKEN" // TODO THIS PROPERLY PLEASE

type EkadashiServer struct {
	db     *database.Service
	cookie map[string]string
}

func InitTelegramBot(service *database.Service) {
	cookie := make(map[string]string)
	s := &EkadashiServer{
		db:     service,
		cookie: cookie,
	}
	token := os.Getenv(ekadashiToken)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	s.ResponseEkadashiBot(bot, &u)
}
