package payload

type Balance struct {
	AcctId string `json:"acctId"`
}

func NewQueryBalancePayload(acctId string) *Balance {
	return &Balance{AcctId: acctId}
}
