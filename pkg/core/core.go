package core

import (
	"fmt"

	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
	"github.com/Feglawy/wetel-cli/pkg/api"
)

type Core struct {
	API api.APIHandler
}

func NewCore(api api.APIHandler) *Core {
	return &Core{
		API: api,
	}
}

func (core *Core) Login(serviceNumber string, password string) error {
	authPayload := *payload.NewAuthPayload(serviceNumber, password)
	err := core.API.Auth(authPayload)
	if err != nil {
		return err
	}
	return nil
}

func (core *Core) GetBalance(accountId string) (float64, error) {
	balancePayload := *payload.NewQueryBalancePayload(accountId)
	balance, err := core.API.QueryBalance(balancePayload)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (core *Core) GetPlans(subscriberId string) (*models.Plan, error) {
	planPayload := *payload.NewQueryFreeUnitPayload(subscriberId)
	return core.API.QueryFreeUnit(planPayload)
}

func (core *Core) getMainOffer(servNumber string) (*models.Offering, error) {
	subOfferpayload := *payload.NewSubOfferPayload(servNumber)
	offers, _ := core.API.GetSubscribedOfferings(subOfferpayload)
	var mainOffer models.Offering
	for _, offer := range offers.OfferingList {
		if offer.Main {
			mainOffer = offer
		}
	}
	return &mainOffer, nil
}

func (core *Core) GetAddonOffers(servNumber string) (models.AddOnOffers, error) {
	payload := *payload.NewGetAddonsPayload(servNumber)
	return core.API.QueryAvailableAddons(payload)
}

func (core *Core) RenewMainOffer(servNumber string, subscriberId string) (string, error) {
	mainOffer, err := core.getMainOffer(servNumber)
	if err != nil {
		return "", fmt.Errorf("couldn't load main offer err: %v", err)
	}
	orderPayload := *payload.NewRenewOfferPayload(subscriberId, *mainOffer)
	bodyStr, err := core.API.OrderCommitForSupplementaryOffer(orderPayload)
	return bodyStr, err
}

func (core *Core) SubscribeToAPlan(subscriberId string, offer models.AddOnOffer) (string, error) {
	o := models.Offering{
		OfferID:     offer.ID,
		OfferEnName: offer.Name,
	}
	subPayload := *payload.NewSubscribeToOfferPayload(subscriberId, o)
	return core.API.OrderCommitForSupplementaryOffer(subPayload)
}
