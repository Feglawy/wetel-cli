package payload

import "github.com/Feglawy/wetel-cli/internal/models"

const (
	SUBSCRIBE   string = "1"
	UNSUBSCRIBE string = "3"
)

type OrderCommitForSupplementaryOffer struct {
	BusinessCode        string                `json:"businessCode"`
	SubscriberID        string                `json:"subscriberId"`
	OwnerType           string                `json:"ownerType"`
	OrderInfo           interface{}           `json:"orderInfo"`
	SubsProductBeanList []subsProductBeanList `json:"subsProductBeanList"`
}

type subsProductBeanList struct {
	SubscriberID   string  `json:"subscriberId"`
	OfferingID     string  `json:"offeringId"`
	OfferingInstID *string `json:"offeringInstId,omitempty"`
	ProdName       string  `json:"prodName"`
	OperType       string  `json:"operType"`
	EffectType     string  `json:"effectType"`
	Quantity       int64   `json:"quantity"`
	PrimaryFlag    string  `json:"primaryFlag"`
}

func newSubsProductBeanItem(subscriberId string, offer models.Offering, operType string) subsProductBeanList {
	s := subsProductBeanList{
		SubscriberID: subscriberId,
		OfferingID:   offer.OfferID,
		ProdName:     "",
		OperType:     operType,
		EffectType:   "2",
		Quantity:     1,
		PrimaryFlag:  "N",
	}
	if operType == UNSUBSCRIBE {
		s.OfferingInstID = &offer.OfferingInstID
	}
	return s
}

func NewRenewOfferPayload(subscriberId string, offer models.Offering) *OrderCommitForSupplementaryOffer {
	return &OrderCommitForSupplementaryOffer{
		BusinessCode: "ChangeOffering",
		SubscriberID: subscriberId,
		OwnerType:    "S",
		OrderInfo:    nil,
		SubsProductBeanList: []subsProductBeanList{
			newSubsProductBeanItem(subscriberId, offer, SUBSCRIBE),
			newSubsProductBeanItem(subscriberId, offer, UNSUBSCRIBE),
		},
	}
}

func NewSubscribeToOfferPayload(subscriberId string, offer models.Offering) *OrderCommitForSupplementaryOffer {
	return &OrderCommitForSupplementaryOffer{
		BusinessCode: "ChangeOffering",
		SubscriberID: subscriberId,
		OwnerType:    "S",
		OrderInfo:    nil,
		SubsProductBeanList: []subsProductBeanList{
			newSubsProductBeanItem(subscriberId, offer, SUBSCRIBE),
		},
	}
}
