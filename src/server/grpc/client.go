package grpc

import (
	"fmt"

	api "github.com/Lighty0410/telegram-bot/src/server/grpc/api"
	"google.golang.org/grpc"
)

// NewGrpcClient creates a connection to the gRPC server.
// If connection succeed returns EkadashiCient interface and nil.
func NewGrpcClient(address string) (api.EkadashiClient, error) {
	conn, err := grpc.Dial("localhost:"+address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to gRPC server: %v", err)
	}
	client := api.NewEkadashiClient(conn)
	return client, nil
}
