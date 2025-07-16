package payload

type Plan struct {
	SubscriberId   string `json:"subscriberId"`
	NeedQueryPoint bool   `json:"needQueryPoint"`
}

func NewQueryFreeUnitPayload(subscriberId string) *Plan {
	return &Plan{
		SubscriberId:   subscriberId,
		NeedQueryPoint: true, // idk
	}
}
