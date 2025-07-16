package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/app"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
)

func QueryBalance(a *app.App, payload payload.Balance) (float64, error) {
	resp, err := a.HandleRequest(http.MethodPost, config.BALANCE_URL, payload)
	if err != nil {
		return 0, fmt.Errorf("couldn't get balance: %v", err)
	}
	var balance models.Balance
	_, err = a.HandleResponse(resp, &balance)
	if err != nil {
		return 0, fmt.Errorf("couldn't get balance err: %v", err)
	}

	return balance.TotalAmount, nil
}
