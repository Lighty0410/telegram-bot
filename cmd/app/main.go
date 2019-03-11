package main

import (
	"log"

	"github.com/Lighty0410/telegram-bot/src/memdb"
	"github.com/Lighty0410/telegram-bot/src/server"
)

func main() {
	service := memdb.NewService()
	err := server.InitTelegramBot(service)
	if err != nil {
		log.Fatal(err)
	}
}
