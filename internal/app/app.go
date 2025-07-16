package app

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/Feglawy/wetel-cli/internal/models"
)

type App struct {
	Client    *http.Client
	UserInfo  models.User
	CSRFToken string
}

func NewApp() *App {
	jar, _ := cookiejar.New(nil)
	return &App{
		Client: &http.Client{Jar: jar},
	}
}
