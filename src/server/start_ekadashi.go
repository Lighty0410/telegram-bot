package server

import (
	"github.com/Lighty0410/telegram-bot/src/database"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)


type EkadashiServer struct {
	db *database.Service
	cookie map[string]string
}

func InitTelegramBot(service *database.Service){
	cookie := make(map[string]string)
	s := &EkadashiServer{
		db: service,
		cookie:cookie,
	}
	bot, err := tgbotapi.NewBotAPI()
	if err != nil {
		log.Println(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	s.ResponseEkadashiBot(bot, &u)
}
