package app

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/Feglawy/wetel-cli/internal/models"
)

type Client struct {
	Client    *http.Client
	UserInfo  models.User
	CSRFToken string
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		Client: &http.Client{Jar: jar},
	}
}
