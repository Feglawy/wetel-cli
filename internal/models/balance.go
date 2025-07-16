package models

import "github.com/tidwall/gjson"

type Balance struct {
	TotalAmount float64
}

func (b *Balance) ScanJson(jsonStr string) error {
	b.TotalAmount = gjson.Get(jsonStr, "body.balanceInfo.0.totalAmount").Float()
	b.TotalAmount /= 10_000
	return nil
}
