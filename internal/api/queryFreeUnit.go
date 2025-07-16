package api

import (
	"fmt"
	"net/http"

	"github.com/Feglawy/wetel-cli/config"
	"github.com/Feglawy/wetel-cli/internal/app"
	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/internal/payload"
)

// This gets your plan's info Qouta remaining gbs etc idk why they named it this way
func QueryFreeUnit(app *app.App, payload payload.Plan) (*models.Plan, error) {
	resp, err := app.HandleRequest(http.MethodPost, config.QUERY_FREE_UNIT, payload)
	if err != nil {
		return nil, fmt.Errorf("plans info request err: %v", err)
	}

	var PlansInfo models.Plan
	_, err = app.HandleResponse(resp, &PlansInfo)
	if err != nil {
		return nil, fmt.Errorf("plans info response err: %v", err)
	}
	return &PlansInfo, nil
}
