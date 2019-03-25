package http

import (
	"fmt"
	"os"

	"github.com/Lighty0410/telegram-bot/src/server/controller"
)

const ekadashiURL = "EKADASHI_URL"

type HttpService struct {
	controller      *controller.Controller
	serverURL       string
	registerURL     string
	loginURL        string
	showEkadashiURL string
}

func NewHttpService(c *controller.Controller) (*HttpService, error) {
	ekadashiURL := os.Getenv(ekadashiURL)
	if ekadashiURL == "" {
		return nil, fmt.Errorf("server URL cannot be empty")
	}
	s := &HttpService{
		controller:      c,
		serverURL:       ekadashiURL,
		registerURL:     "/register",
		loginURL:        "/login",
		showEkadashiURL: "/ekadashi/next",
	}
	return s, nil
}
