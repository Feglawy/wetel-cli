package models

import (
	"github.com/tidwall/gjson"
)

type AddOnOffer struct {
	ID                string
	Name              string
	AlreadySubscribed bool
	Price             float64
	GroupName         string
	GroupType         string
}

type AddOnOffers []AddOnOffer

func (a *AddOnOffers) ScanJson(jsonStr string) error {
	root := gjson.Get(jsonStr, "body.availableAddOnOfferingsList")

	// recursive function
	var parseGroup func(group gjson.Result, parentGroup string)

	parseGroup = func(group gjson.Result, parentGroup string) {
		groupName := group.Get("groupEnName").String()

		// Check if current group has offers
		if offers := group.Get("availableAddOnOfferingList"); offers.Exists() {
			offers.ForEach(func(_, offer gjson.Result) bool {
				*a = append(*a, AddOnOffer{
					ID:                offer.Get("offerId").String(),
					Name:              offer.Get("offerEnName").String(),
					AlreadySubscribed: offer.Get("alreadySubscribed").Bool(),
					Price:             offer.Get("price").Float(),
					GroupName:         parentGroup,
					GroupType:         groupName,
				})
				return true
			})
		}

		// Check if it has child groups and parse them
		if children := group.Get("childGroups"); children.Exists() {
			children.ForEach(func(_, child gjson.Result) bool {
				parseGroup(child, groupName)
				return true
			})
		}
	}

	// Start parsing top-level groups
	root.ForEach(func(_, topGroup gjson.Result) bool {
		parseGroup(topGroup, "")
		return true
	})

	return nil
}
