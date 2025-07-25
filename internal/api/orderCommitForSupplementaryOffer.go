package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/payload"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/utils"
	"github.com/tidwall/gjson"
)

func OrderCommitForSupplementaryOffer(app *app.Client, payload payload.OrderCommitForSupplementaryOffer) (string, error) {
	resp, err := app.HandleRequest(http.MethodPost, config.ORDER_SUPPLEMENTARY_OFFER, payload)
	if err != nil {
		return "", fmt.Errorf("error with request order supplementry offer err : %v", err)
	}
	body, err := app.HandleResponse(resp, nil)
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
