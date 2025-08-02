package payload

type Addons struct {
	AccessNumber      string `json:"accessNumber"`
	BusinessCode      string `json:"businessCode"`
	ChannelID         string `json:"channelId"`
	CurrentMainID     string `json:"currentMainId"`
	GroupID           string `json:"groupId"`
	Locale            string `json:"locale"`
	NumberServiceType string `json:"numberServiceType"`
	PaidType          string `json:"paidType"`
}

func NewGetAddonsPayload(srvNumber string) *Addons {
	return &Addons{
		AccessNumber:      srvNumber,
		BusinessCode:      "",
		ChannelID:         "WEB_APP",
		CurrentMainID:     "",
		GroupID:           "",
		Locale:            "en-US",
		NumberServiceType: "FBB",
		PaidType:          "",
	}
}
