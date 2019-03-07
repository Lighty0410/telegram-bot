package main

import (
	"github.com/Lighty0410/telegram-bot/src/database"
	"github.com/Lighty0410/telegram-bot/src/server"
)

func main() {
	service := database.NewService()
	server.InitTelegramBot(service)
}
