package menu

import (
	"fmt"

	"github.com/Feglawy/wetel-cli/internal/models"
	"github.com/Feglawy/wetel-cli/pkg/app"
	"github.com/Feglawy/wetel-cli/pkg/core"
)

func Menu(app *app.Client) {
	balance, err := core.GetBalance(app)
	if err != nil {
		fmt.Printf("Error fetching balance: %s\n", err)
		return
	}

	plan, err := core.GetPlans(app)
	if err != nil {
		fmt.Printf("Error fetching plans: %s\n", err)
		return
	}

	overview(&app.UserInfo, balance, plan)

	for {
		fmt.Print(`
______________________________
Menu:
1. Show detailed plan overview
2. Renew main plan
3. Refresh
4. Exit
>> `)
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			offerUsageOverview(plan)
		case 2:
			body, err := core.RenewMainOffer(app)
			if err != nil {
				fmt.Printf("Failed to renew: %s\n", err)
			} else {
				fmt.Println(body)
			}
		case 3:
			Menu(app)
			return
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown option. Please try again.")
		}
	}
}

func overview(user *models.User, balance float64, plan *models.Plan) {
	fmt.Printf("\nHello, %s\n", user.CustomerName)
	fmt.Printf("Current balance: %.2f\n", balance)
	fmt.Println("______________________________")
	fmt.Println("Usage overview")
	fmt.Printf("Used: %.2f GB\n", plan.Used())
	fmt.Printf("Remaining: %.2f GB\n", plan.Remain())
}

func offerUsageOverview(plan *models.Plan) {
	fmt.Println("______________________________")
	fmt.Println("Offers usage overview")
	for _, offer := range plan.Offers {
		fmt.Printf("Offer: %s\n", offer.Name)
		fmt.Printf("Total: %.2f GB\n", offer.InitialAmount)
		fmt.Printf("Remaining: %.2f GB\n", offer.CurrentAmount)
		fmt.Printf("Used: %.2f GB\n", offer.InitialAmount-offer.CurrentAmount)
		fmt.Println("---------------------------")
	}
}
