package server

import (
	"fmt"
	"os"

	"github.com/Lighty0410/telegram-bot/src/server/grpc"

	"github.com/Lighty0410/telegram-bot/src/server/controller"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const ekadashiURL = "EKADASHI_URL"
const ekadashiToken = "EKADASHI_TOKEN"

// EkadashiBot contains information about database and URLs.
type EkadashiBot struct {
	grpc            *grpc.GrpcService
	controller      controller.Controller
	serverURL       string
	registerURL     string
	loginURL        string
	showEkadashiURL string
}

// InitTelegramBot initialize Telegram-bot.
// It also have a basic structures for token and URLs setup.
func InitTelegramBot(gRPC *grpc.GrpcService) error {
	ekadashiURL := os.Getenv(ekadashiURL)
	if ekadashiURL == "" {
		return fmt.Errorf("server URL cannot be empty")
	}
	s := &EkadashiBot{
		grpc:            gRPC,
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
