package main

import (
	"log"

	"github.com/Lighty0410/telegram-bot/src/memdb"
	"github.com/Lighty0410/telegram-bot/src/server"
	"github.com/Lighty0410/telegram-bot/src/server/controller"
	"github.com/Lighty0410/telegram-bot/src/server/grpc"
)

func main() {
	memdb := memdb.NewService()
	controller := controller.NewController(memdb)
	gRPC, err := grpc.NewGrpcClient("50051", controller)
	if err != nil {
		log.Fatalf("cannot get gRPC client: %v", err)
	}
	err = server.InitTelegramBot(gRPC)
	if err != nil {
		log.Fatal(err)
	}
}
