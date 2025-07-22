package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/payload"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/tidwall/gjson"
)

func Auth(a *app.Client, payload payload.AuthPayload) error {
	resp, err := a.HandleRequest(http.MethodPost, config.AUTH_URL, payload)
	if err != nil {
		return fmt.Errorf("couldn't login err: %v", err)
	}
	body, err := a.HandleResponse(resp, &a.UserInfo)
	if err != nil {
		return fmt.Errorf("couldn't login err: %v", err)
	}

	a.CSRFToken = gjson.Get(body, "body.token").String()

	return nil
}
