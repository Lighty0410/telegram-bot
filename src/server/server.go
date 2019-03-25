package server

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// ResponseEkadashiBot is major method for this microservice.
// Basically it's a go-routine that retrieves user's query and handle it.
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
			err := s.grpc.Register(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println("cannot register user by gRPC method: ", err)
				err := s.http.Register(strconv.FormatInt(update.Message.Chat.ID, 10))
				if err != nil {
					log.Println("cannot register user by HTTP method: ", err)
				}
			}
		case "login":
			err := s.grpc.HandleLogin(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println("cannot login user by gRPC method: ", err)
				err = s.http.Login(strconv.FormatInt(update.Message.Chat.ID, 10))
				if err != nil {
					log.Println("cannot login user by HTTP method: ", err)
				}
			}
		case "ekadashi":
			ekadashiDate, err := s.grpc.ShowEkadashi(strconv.FormatInt(update.Message.Chat.ID, 10))
			if err != nil {
				log.Println("cannot get ekadahi date by gRPC method: ", err)
				ekadashiDate, err := s.http.ShowEkadashi(strconv.FormatInt(update.Message.Chat.ID, 10))
				if err != nil {
					log.Println("cannot get ekadashi date by HTTP method: ", err)
					continue
				}
				update.Message.Text = ekadashiDate
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				_, err = bot.Send(msg)
				if err != nil {
					log.Println("cannot send message to the telegram: ", err)
					continue
				}
				continue
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
