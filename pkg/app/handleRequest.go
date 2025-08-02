package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Feglawy/wetel-cli/utils"
)

func (a *Client) HandleRequest(method, url string, payload any) (*http.Response, error) {
	var body io.Reader

	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("marshal payload: %w", err)
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	utils.SetHeaders(req, a.csrfToken)

	return a.Client.Do(req)
}
