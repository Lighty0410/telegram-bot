package server

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *EkadashiBot) ResponseEkadashiBot(bot *tgbotapi.BotAPI, u tgbotapi.UpdateConfig) {
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println("cannot get updates from channel: ", err)
		return
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "start":
			err := s.handleRegistration(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println("cannot register user: ", err)
			}
		case "login":
			err := s.handleLogin(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println("cannot login user: ", err)
			}
		case "ekadashi":
			ekadashiDate, err := s.showEkadashiHandler(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println(err)
			}
			update.Message.Text = ekadashiDate
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			_, err = bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
