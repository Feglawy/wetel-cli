package core

import (
	"fmt"

	"github.com/Feglawy/wetel-cli/internal/api"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
	"github.com/Feglawy/wetel-cli/pkg/app"
)

func Login(serviceNumber string, password string, a *app.Client) error {
	authPayload := *payload.NewAuthPayload(serviceNumber, password)
	err := api.Auth(a, authPayload)
	if err != nil {
		return err
	}
	return nil
}

func GetBalance(a *app.Client) (float64, error) {
	balancePayload := *payload.NewQueryBalancePayload(a.UserInfo.AccountId)
	balance, err := api.QueryBalance(a, balancePayload)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func GetPlans(a *app.Client) (*models.Plan, error) {
	planPayload := *payload.NewQueryFreeUnitPayload(a.UserInfo.SubscriberId)
	plans, err := api.QueryFreeUnit(a, planPayload)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func getMainOffer(a *app.Client) (*models.Offering, error) {
	subOfferpayload := *payload.NewSubOfferPayload(a.UserInfo.ServNumber)
	offers, _ := api.GetSubscribedOfferings(a, subOfferpayload)
	var mainOffer models.Offering
	for _, offer := range offers.OfferingList {
		if offer.Main {
			mainOffer = offer
		}
	}
	return &mainOffer, nil
}

func RenewMainOffer(a *app.Client) (string, error) {
	mainOffer, err := getMainOffer(a)
	if err != nil {
		return "", fmt.Errorf("couldn't load main offer err: %v", err)
	}
	orderPayload := *payload.NewRenewOfferPayload(a.UserInfo.SubscriberId, *mainOffer)
	bodyStr, err := api.OrderCommitForSupplementaryOffer(a, orderPayload)
	return bodyStr, err
}
