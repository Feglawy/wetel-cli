package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/app"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
)

func GetSubscribedOfferings(a *app.App, payload payload.SubOffer) (*models.SubOffers, error) {
	resp, err := a.HandleRequest(http.MethodPost, config.SUBSCRIBED_OFFERS_DATA, payload)
	if err != nil {
		return nil, fmt.Errorf("error with request subscribed offers err: %v", err)
	}
	var offers models.SubOffers
	_, err = a.HandleResponse(resp, &offers)
	if err != nil {
		return nil, fmt.Errorf("error with response subscribed offers err: %v", err)
	}
	return &offers, nil
}
