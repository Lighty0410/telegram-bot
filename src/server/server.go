package server

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (s *EkadashiServer) ResponseEkadashiBot(bot *tgbotapi.BotAPI, u *tgbotapi.UpdateConfig) {
	updates, err := bot.GetUpdatesChan(*u)
	if err != nil {
		log.Println("cannot get updates from channel: ", err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "register":
			err := s.handleRegistration(strconv.FormatInt(int64(update.Message.Chat.ID), 10))
			if err != nil {
				log.Println("cannot register user: ", err)
			}
		case "login":
			err := s.handleLogin(strconv.FormatInt(int64(update.Message.Chat.ID), 10))
			if err != nil {
				log.Println("cannot login user: ", err)
			}
		case "showEkadashi":
			ekadashiDate, err := s.showEkadashiHandler(strconv.FormatInt(int64(update.Message.Chat.ID), 10))
			if err != nil {
				log.Println(err)
			}
			update.Message.Text = ekadashiDate
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}
}
