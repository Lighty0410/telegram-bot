package main

import (
	"log"

	"github.com/Lighty0410/telegram-bot/src/memdb"
	"github.com/Lighty0410/telegram-bot/src/server"
	"github.com/Lighty0410/telegram-bot/src/server/controller"
	"github.com/Lighty0410/telegram-bot/src/server/grpc"
	"github.com/Lighty0410/telegram-bot/src/server/http"
)

func main() {
	memdb := memdb.NewService()
	controller := controller.NewController(memdb)
	gRPC, err := grpc.NewGrpcClient("90000", controller)
	if err != nil {
		log.Fatalf("cannot get gRPC client: %v", err)
	}
	http, err := http.NewHttpService(controller)
	err = server.InitTelegramBot(gRPC, http)
	if err != nil {
		log.Fatal(err)
	}
}
