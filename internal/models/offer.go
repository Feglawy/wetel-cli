package models

import "github.com/tidwall/gjson"

type Offer struct {
	Type          string
	Name          string
	InitialAmount float64
	CurrentAmount float64
	RemainingDays int
}

/*
Type = `body.freeUnitBeanDetailList.originType`
Name = `body.freeUnitBeanDetailList.offeringName`
InitialAmount = `body.freeUnitBeanDetailList.initialAmount`
CurrentAmount = `body.freeUnitBeanDetailList.currentAmount`
RemainingDays = `body.freeUnitBeanDetailList.remainingDaysForRenewal`
*/

func (o *Offer) ScanJson(jsonStr string) error {
	o.Type = gjson.Get(jsonStr, "body.freeUnitBeanDetailList.0.originType").String()
	o.Name = gjson.Get(jsonStr, "body.freeUnitBeanDetailList.0.offeringName").String()
	o.InitialAmount = gjson.Get(jsonStr, "body.freeUnitBeanDetailList.0.initialAmount").Float()
	o.CurrentAmount = gjson.Get(jsonStr, "body.freeUnitBeanDetailList.0.currentAmount").Float()
	d := gjson.Get(jsonStr, "body.freeUnitBeanDetailList.0.remainingDaysForRenewal").Int()
	o.RemainingDays = int(d)
	return nil
}
