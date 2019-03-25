package server

import (
	"context"
	"fmt"
	"time"

	api "github.com/Lighty0410/telegram-bot/src/server/grpc/api"
)

// EkadashiDate contains information about ekadashi date.
type EkadashiDate struct {
	Date string `json:"date"`
}

// showEkadashiHandler shows next ekadashi day based on another server.
func (s *EkadashiBot) showEkadashiHandler(username string) (string, error) {
	user, err := s.controller.GetUser(username)
	if err != nil {
		return "", fmt.Errorf("cannot get user :%v", err)
	}
	date, err := s.client.ShowEkadashi(context.Background(), &api.ShowEkadashiRequest{
		Session: &api.Session{Token: user.Token},
	})
	if err != nil {
		return "", fmt.Errorf("cannot get ekadashi date: %v", err)
	}
	ekadashi := time.Unix(date.Ekadashi, 0)
	return ekadashi.Format("January 2 2006"), nil
}
