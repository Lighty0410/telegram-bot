package main

import (
	"log"

	"github.com/Lighty0410/telegram-bot/src/memdb"
	"github.com/Lighty0410/telegram-bot/src/server"
)

func main() {
	memdb := memdb.NewService()
	err := server.InitTelegramBot(memdb)
	if err != nil {
		log.Fatal(err)
	}
}
