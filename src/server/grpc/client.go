package grpc

import (
	"fmt"

	"github.com/Lighty0410/telegram-bot/src/server/controller"

	api "github.com/Lighty0410/telegram-bot/src/server/grpc/api"
	"google.golang.org/grpc"
)

type GrpcService struct {
	client     api.EkadashiClient
	controller *controller.Controller
}

// NewGrpcClient creates a connection to the gRPC server.
// If connection succeed returns EkadashiCient interface and nil.
func NewGrpcClient(address string, controller *controller.Controller) (*GrpcService, error) {
	conn, err := grpc.Dial("localhost:"+address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to gRPC server: %v", err)
	}
	client := api.NewEkadashiClient(conn)
	s := &GrpcService{
		client:     client,
		controller: controller,
	}
	return s, nil
}
