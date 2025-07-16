package models

import (
	"github.com/tidwall/gjson"
)

type User struct {
	SubscriberId string
	CustId       string
	AccountId    string
	CustomerName string
	ServNumber   string
}

/*
`body.subscriber.subscriberId`
`body.subscriber.custId`
`body.subscriber.accountId`
`body.customer.custName`
`body.subscriber.servNumber`
*/

func (u *User) ScanJson(jsonStr string) error {
	u.SubscriberId = gjson.Get(jsonStr, "body.subscriber.subscriberId").String()
	u.CustId = gjson.Get(jsonStr, "body.subscriber.custId").String()
	u.AccountId = gjson.Get(jsonStr, "body.subscriber.accountId").String()
	u.CustomerName = gjson.Get(jsonStr, "body.customer.custName").String()
	u.ServNumber = gjson.Get(jsonStr, "body.subscriber.servNumber").String()
	return nil
}
