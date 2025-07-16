package models

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type SubOffers struct {
	OfferingList []Offering `json:"offeringList"`
}

type Offering struct {
	OfferID                 string `json:"offerId"`
	OfferingInstID          string `json:"offeringInstId"`
	OfferEnName             string `json:"offerEnName"`
	RemainingDaysForRenewal *int64 `json:"remainingDaysForRenewal,omitempty"`
	SubscriptionDays        *int64 `json:"subscriptionDays,omitempty"`
	Primary                 bool   `json:"primary"`
	Main                    bool   `json:"main"`
	AddOn                   bool   `json:"addOn"`
	Price                   int64  `json:"price"`
	Renewable               *bool  `json:"renewable,omitempty"`
}

func (s *SubOffers) ScanJson(jsonStr string) error {
	body := gjson.Get(jsonStr, "body").String()
	return json.Unmarshal([]byte(body), s)
}
