package server

import (
	"fmt"
	"os"

	"github.com/Lighty0410/telegram-bot/src/memdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const ekadashiURL = "EKADASHI_URL"
const ekadashiToken = "EKADASHI_TOKEN"

type EkadashiBot struct {
	db              *memdb.Service
	serverURL       string
	registerURL     string
	loginURL        string
	showEkadashiURL string
}

func InitTelegramBot(service *memdb.Service) error {
	ekadashiURL := os.Getenv(ekadashiURL)
	if ekadashiURL == "" {
		return fmt.Errorf("server URL cannot be empty")
	}
	s := &EkadashiBot{
		db:              service,
		serverURL:       ekadashiURL,
		registerURL:     "/register",
		loginURL:        "/login",
		showEkadashiURL: "/ekadashi/next",
	}
	token := os.Getenv(ekadashiToken)
	if token == "" {
		return fmt.Errorf("ekadashi token cannot be empty")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	s.ResponseEkadashiBot(bot, u)
	return nil
}
