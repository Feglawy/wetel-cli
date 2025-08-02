package ui

import (
	"fmt"
	"strings"

	"github.com/Feglawy/wetel-cli/internal/models"
)

func Overview(user *models.User, balance float64, plan *models.Plan) {
	fmt.Printf("\nHello, %s\n", user.CustomerName)
	fmt.Printf("Current balance: %.2f\n", balance)
	fmt.Println("______________________________")
	fmt.Println("Usage overview")
	fmt.Printf("Used: %.2f GB\n", plan.Used())
	fmt.Printf("Remaining: %.2f GB\n", plan.Remain())
	fmt.Printf("Remaining days: %v\n", plan.GetRemainingDays())
	fmt.Println("______________________________")
}

func OfferUsageOverview(plan *models.Plan) {
	fmt.Println("Offers usage overview")
	for _, offer := range plan.Offers {
		fmt.Println("______________________________")
		fmt.Printf("Offer: %s\n", offer.Name)
		fmt.Printf("Remaining days: %v\n", offer.RemainingDays)
		fmt.Printf("Total: %.2f GB\n", offer.InitialAmount)
		fmt.Printf("Used: %.2f GB\n", offer.InitialAmount-offer.CurrentAmount)
		fmt.Printf("Remaining: %.2f GB\n", offer.CurrentAmount)
	}
	fmt.Println("______________________________")
}

func ChooseAnADDON(offers models.AddOnOffers) *models.AddOnOffer {
	for i, offer := range offers {
		fmt.Printf("%v - type: %v - offer: %v - price: %v \n", i+1, offer.GroupType, offer.Name, offer.Price)
	}
	var chosenOption int
	for {
		fmt.Print("Enter an option: ")
		fmt.Scan(&chosenOption)
		chosenOption -= 1
		if chosenOption < len(offers) && chosenOption >= 0 {
			break
		}
		fmt.Println("invalid option please try again")
	}

	chosenOffer := offers[chosenOption]
	fmt.Println("You have choosed")
	fmt.Printf("type: %v - offer: %v - price: %v \n", chosenOffer.GroupType, chosenOffer.Name, chosenOffer.Price)
	fmt.Println("Are you sure (y/n): ")
	var sure string
	fmt.Scan(&sure)
	sure = strings.ToLower(sure)
	if sure == "y" {
		return &chosenOffer
	}
	return nil
}
