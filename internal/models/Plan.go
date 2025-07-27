package models

import "github.com/tidwall/gjson"

type Plan struct {
	Offers []Offer
}

func (p *Plan) Total() float64 {
	var total float64
	for _, o := range p.Offers {
		total += o.InitialAmount
	}
	return total
}

func (p *Plan) GetRemainingDays() int {

	for _, offer := range p.Offers {
		if offer.Type == "Main Quota" {
			return offer.RemainingDays
		}
	}
	return 0
}

func (p *Plan) Used() float64 {
	var used float64
	for _, o := range p.Offers {
		used += o.InitialAmount - o.CurrentAmount
	}
	return used
}

func (p *Plan) Remain() float64 {
	return p.Total() - p.Used()
}

func (p *Plan) ScanJson(jsonStr string) error {
	offers := gjson.Get(jsonStr, "body.0.freeUnitBeanDetailList")
	p.Offers = make([]Offer, 0)

	offers.ForEach(func(_, value gjson.Result) bool {
		var offer Offer
		offer.Type = value.Get("originType").String()
		offer.Name = value.Get("offeringName").String()
		offer.InitialAmount = value.Get("initialAmount").Float()
		offer.CurrentAmount = value.Get("currentAmount").Float()
		offer.RemainingDays = int(value.Get("remainingDaysForRenewal").Int())
		p.Offers = append(p.Offers, offer)
		return true
	})
	return nil
}
