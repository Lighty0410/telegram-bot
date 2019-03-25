package server

import (
	"fmt"
	"os"

	"github.com/Lighty0410/telegram-bot/src/server/http"

	"github.com/Lighty0410/telegram-bot/src/server/grpc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const ekadashiToken = "EKADASHI_TOKEN"

// EkadashiBot contains information about database and URLs.
type EkadashiBot struct {
	grpc *grpc.Service
	http *http.Service
}

// InitTelegramBot initialize Telegram-bot.
// It also have a basic structures for token and URLs setup.
func InitTelegramBot(gRPC *grpc.Service, http *http.Service) error {
	s := &EkadashiBot{
		grpc: gRPC,
		http: http,
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
