package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/utils"
	"github.com/tidwall/gjson"
)

type APIHandler interface {
	Auth(payload payload.AuthPayload) error
	QueryBalance(payload payload.Balance) (float64, error)
	QueryFreeUnit(payload payload.Plan) (*models.Plan, error)
	GetSubscribedOfferings(payload payload.SubOffer) (*models.SubOffers, error)
	OrderCommitForSupplementaryOffer(payload payload.OrderCommitForSupplementaryOffer) (string, error)
}

type API struct {
	App app.ClientHandler
}

func NewAPI(app app.ClientHandler) *API {
	return &API{
		App: app,
	}
}

func (api API) Auth(payload payload.AuthPayload) error {
	resp, err := api.App.HandleRequest(http.MethodPost, config.AUTH_URL, payload)
	if err != nil {
		return fmt.Errorf("couldn't login err: %v", err)
	}
	var userInfo models.User
	body, err := api.App.HandleResponse(resp, &userInfo)
	if err != nil {
		return fmt.Errorf("couldn't login err: %v", err)
	}
	api.App.SetUserInfo(userInfo)
	api.App.SetCSRFToken(gjson.Get(body, "body.token").String())

	return nil
}

func (api API) GetSubscribedOfferings(payload payload.SubOffer) (*models.SubOffers, error) {
	resp, err := api.App.HandleRequest(http.MethodPost, config.SUBSCRIBED_OFFERS_DATA, payload)
	if err != nil {
		return nil, fmt.Errorf("error with request subscribed offers err: %v", err)
	}
	var offers models.SubOffers
	_, err = api.App.HandleResponse(resp, &offers)
	if err != nil {
		return nil, fmt.Errorf("error with response subscribed offers err: %v", err)
	}
	return &offers, nil
}

func (api API) OrderCommitForSupplementaryOffer(payload payload.OrderCommitForSupplementaryOffer) (string, error) {
	resp, err := api.App.HandleRequest(http.MethodPost, config.ORDER_SUPPLEMENTARY_OFFER, payload)
	if err != nil {
		return "", fmt.Errorf("error with request order supplementry offer err : %v", err)
	}
	body, err := api.App.HandleResponse(resp, nil)
	body = utils.GetIndentedJson([]byte(body))
	if err != nil {
		return body, fmt.Errorf("error with response order supplementry offer err : %v", err)
	}
	retMsg := ""
	if utils.IsRespSuccessful(body) {
		retMsg = gjson.Get(body, "body.retMsg").String()
	}
	return retMsg, nil
}

func (api API) QueryBalance(payload payload.Balance) (float64, error) {
	resp, err := api.App.HandleRequest(http.MethodPost, config.BALANCE_URL, payload)
	if err != nil {
		return 0, fmt.Errorf("couldn't get balance: %v", err)
	}
	var balance models.Balance
	_, err = api.App.HandleResponse(resp, &balance)
	if err != nil {
		return 0, fmt.Errorf("couldn't get balance err: %v", err)
	}

	return balance.TotalAmount, nil
}

// This gets your plan's info Qouta remaining gbs etc idk why they named it this way
func (api API) QueryFreeUnit(payload payload.Plan) (*models.Plan, error) {
	resp, err := api.App.HandleRequest(http.MethodPost, config.QUERY_FREE_UNIT, payload)
	if err != nil {
		return nil, fmt.Errorf("plans info request err: %v", err)
	}

	var PlansInfo models.Plan
	_, err = api.App.HandleResponse(resp, &PlansInfo)
	if err != nil {
		return nil, fmt.Errorf("plans info response err: %v", err)
	}
	return &PlansInfo, nil
}

func (api API) QueryAvailableAddons(payload payload.Addons) (models.AddOnOffers, error) {
	resp, err := api.App.HandleRequest(http.MethodPost, config.QUERY_ADDONS, payload)
	if err != nil {
		return nil, fmt.Errorf("request err: %v", err)
	}

	// handle response
	var offers models.AddOnOffers
	_, err = api.App.HandleResponse(resp, &offers)
	if err != nil {
		return nil, fmt.Errorf("addon offers response err: %v", err)
	}
	return offers, nil
}
