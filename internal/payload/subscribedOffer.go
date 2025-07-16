package payload

type SubOffer struct {
	GroupID           string `json:"groupId"`
	ServiceNumber     string `json:"msisdn"`
	NumberServiceType string `json:"numberServiceType"`
}

func NewSubOfferPayload(serviceNum string) *SubOffer {
	return &SubOffer{
		GroupID:           "",
		ServiceNumber:     serviceNum,
		NumberServiceType: "FBB",
	}
}
