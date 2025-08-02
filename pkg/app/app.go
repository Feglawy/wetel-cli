package app

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/Feglawy/wetel-cli/internal/models"
)

type ClientHandler interface {
	GetUserInfo() models.User
	SetUserInfo(info models.User)
	GetCSRFToken() string
	SetCSRFToken(token string)
	HandleRequest(method string, url string, payload any) (*http.Response, error)
	HandleResponse(resp *http.Response, out models.Scannable) (string, error)
}

type Client struct {
	Client    *http.Client
	userInfo  models.User
	csrfToken string
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		Client: &http.Client{
			Jar:     jar,
			Timeout: 20 * time.Second,
		},
	}
}

func (c *Client) SetUserInfo(info models.User) {
	c.userInfo = info
}

func (c *Client) GetUserInfo() models.User {
	return c.userInfo
}

func (c *Client) SetCSRFToken(token string) {
	c.csrfToken = token
}
func (c *Client) GetCSRFToken() string {
	return c.csrfToken
}
